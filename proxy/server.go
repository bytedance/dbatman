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

// Server is the proxy server. It handle the request from frontend, process and dispatch
// queries, picking right backend conn due to the request context.

type LimitReqNode struct {
	excess     int64
	last       int64
	query      string
	count      int64
	lastSecond int64 //Last second to refresh the excess?
}

type Server struct {
	cfg *config.Conf

	// nodes map[string]*Node

	// schemas map[string]*Schema

	// users    *userAuth
	mu *sync.Mutex

	fingerprints map[string]*LimitReqNode
	listener     net.Listener
	running      bool
}

func NewServer(cfg *config.Conf) (*Server, error) {
	s := new(Server)

	s.cfg = cfg

	var err error

	s.fingerprints = make(map[string]*LimitReqNode)
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

	for s.running {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Warning("accept error %s", err.Error())
			continue
		}

		go s.onConn(conn)
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
		log.Warnf("handshake error: %s", err)
		return
	}

	if err := session.Run(); err != nil {
		// TODO
		// session.WriteError(NewDefaultError(err))
		if err == errSessionQuit {
			return
		}

		log.Warnf("session run error: %s", err.Error())
	}
}
