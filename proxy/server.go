package proxy

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"time"

	"sync"
	"syscall"

	"github.com/bytedance/dbatman/config"
	_ "github.com/bytedance/dbatman/database/mysql"
	"github.com/ngaut/log"
)

var startNum = 0
var closeNum = 0

var sessionChan = make(chan int64, 256)

type LimitReqNode struct {
	excess     int64
	last       int64
	query      string
	count      int64
	lastSecond int64 //Last second to refresh the excess?

	start        int64 //qps start time by millsecond
	lastcount    int64 //last count rep num means qps
	currentcount int64 //repnum in current 1s dperiod
}

type Ip struct {
	ip          string
	mu          sync.Mutex
	printfinger map[string]*LimitReqNode
}
type User struct {
	user   string
	iplist map[string]*Ip
}
type Server struct {
	cfg *config.Conf

	// nodes map[string]*Node

	// schemas map[string]*Schema

	// users    *userAuth
	mu *sync.Mutex
	// users        map[string]*User
	//qps base on fingerprint
	fingerprints map[string]*LimitReqNode
	//qps base on server
	qpsOnServer *LimitReqNode
	listener    net.Listener
	running     bool
	restart     bool
	wg          sync.WaitGroup
}

func NewServer(cfg *config.Conf) (*Server, error) {
	s := new(Server)

	s.cfg = cfg

	var err error

	s.fingerprints = make(map[string]*LimitReqNode)
	// s.users = make(map[string]*User)
	// s.qpsOnServer = &LimitReqNode{}
	s.mu = &sync.Mutex{}
	s.restart = false
	port := s.cfg.GetConfig().Global.Port

	// get listenfd from file when restart
	if os.Getenv("_GRACEFUL_RESTART") == "true" {
		log.Info("graceful restart with previous listenfd")

		//get the linstenfd
		file := os.NewFile(3, "")
		s.listener, err = net.FileListener(file)
		if err != nil {
			log.Warn("get linstener err ")
		}

	} else {
		s.listener, err = net.Listen("tcp4", fmt.Sprintf(":%d", port))
	}
	if err != nil {
		return nil, err
	}

	log.Infof("Dbatman Listen(tcp4) at [%d]", port)
	return s, nil
}

func (s *Server) Serve() error {
	log.Debug("this is ddbatman v3")
	s.running = true
	var sessionId int64 = 0
	for s.running {
		select {
		case sessionChan <- sessionId:
			//do nothing
		default:
			//warnning!
			log.Warnf("TASK_CHANNEL is full!")
		}

		conn, err := s.Accept()
		if err != nil {
			log.Warning("accept error %s", err.Error())
			continue
		}
		//allocate a sessionId for a session
		go s.onConn(conn)
		sessionId += 1
	}
	if s.restart == true {
		log.Debug("Begin to restart graceful")
		listenerFile, err := s.listener.(*net.TCPListener).File()
		if err != nil {
			log.Fatal("Fail to get socket file descriptor:", err)
		}
		listenerFd := listenerFile.Fd()

		os.Setenv("_GRACEFUL_RESTART", "true")
		execSpec := &syscall.ProcAttr{
			Env:   os.Environ(),
			Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd(), listenerFd},
		}
		fork, err := syscall.ForkExec(os.Args[0], os.Args, execSpec)
		if err != nil {
			return fmt.Errorf("failed to forkexec: %v", err)
		}

		log.Infof("start new process success, pid %d.", fork)
	}
	timeout := time.NewTimer(time.Minute)
	wait := make(chan struct{})
	go func() {
		s.wg.Wait()
		wait <- struct{}{}
	}()

	select {
	case <-timeout.C:
		log.Error("server : Waittimeout error when close the service")
		return nil
	case <-wait:
		log.Info("server : all goroutine has been done")
		return nil
	}
	return nil
}
func (s *Server) Accept() (net.Conn, error) {

	conn, err := s.listener.Accept()
	if err != nil {
		return nil, err
	}
	// tc.SetKeepAlive(true)
	// tc.SetKeepAlivePeriod(3 * time.Minute)

	s.wg.Add(1)
	startNum += 1
	// log.Info("wait group add 1 total is :", startNum)

	return conn, nil
}

// TODO check this function if it need routine-safe
func (s *Server) Close() {
	s.running = false
	if s.listener != nil {
		s.listener.Close()
		s.listener = nil
	}
}
func (s *Server) Restart() {
	s.running = false
	s.restart = true
	if s.listener != nil {
		//s.listener.Close()
		//s.listener = nil
	}
}

func (s *Server) onConn(c net.Conn) {
	session := s.newSession(c)

	defer func() {
		if !debug {
			if err := recover(); err != nil {
				const size = 4096
				buf := make([]byte, size)
				buf = buf[:runtime.Stack(buf, false)]
				log.Fatalf("onConn panic %v: %v\n%s", c.RemoteAddr().String(), err, buf)
			}
		}

		session.Close()
	}()
	// Handshake error, here we do not need to close the conn
	if err := session.Handshake(); err != nil {
		log.Warnf("session %d handshake error: %s", session.sessionId, err)
		return
	}

	if err := session.Run(); err != nil {
		// TODO

		// session.WriteError(NewDefaultError(err))
		session.Close()
		if err == errSessionQuit {

			log.Warnf("session %d: %s", session.sessionId, err.Error())
			// return
		}
		closeNum += 1
		s.wg.Done()
		log.Info("wait group add 1 total is :", closeNum)
		log.Warnf("session %d:session run error: %s", session.sessionId, err.Error())
		return
	}
}
