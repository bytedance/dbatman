package proxy

import (
	"fmt"
	"net"
	"runtime"

	"sync"

	"github.com/bytedance/dbatman/config"
	_ "github.com/bytedance/dbatman/database/mysql"
	"github.com/ngaut/log"
)

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
}

func NewServer(cfg *config.Conf) (*Server, error) {
	s := new(Server)

	s.cfg = cfg

	var err error

	s.fingerprints = make(map[string]*LimitReqNode)
	// s.users = make(map[string]*User)
	// s.qpsOnServer = &LimitReqNode{}
	s.mu = &sync.Mutex{}

	port := s.cfg.GetConfig().Global.Port
	s.listener, err = net.Listen("tcp4", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}

	log.Infof("Dbatman Listen(tcp4) at [%d]", port)
	return s, nil
}

func (s *Server) Serve() error {
	s.running = true
	var sessionId int64 = 0
	for s.running {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Warning("accept error %s", err.Error())
			continue
		}
		//allocate a sessionId for a session
		select {
		case sessionChan <- sessionId:
			//do nothing
		default:
			//warnning!
			log.Warnf("TASK_CHANNEL is full!")
		}
		go s.onConn(conn)
		sessionId += 1
	}

	return nil
}

// TODO check this function if it need routine-safe
func (s *Server) Close() {
	s.running = false
	if s.listener != nil {
		s.listener.Close()
		s.listener = nil
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
				log.Fatal("onConn panic %v: %v\n%s", c.RemoteAddr().String(), err, buf)
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
			return
		}
		log.Warnf("session %d:session run error: %s", session.sessionId, err.Error())
	}
}
