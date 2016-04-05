package proxy

import (
	"github.com/bytedance/dbatman/config"
	"github.com/juju/errors"
	"github.com/ngaut/log"
	"net"
	"runtime"
)

// Server is the proxy server. It handle the request from frontend, process and dispatch
// queries, picking right backend conn due to the request context.
type Server struct {
	cfg *config.ProxyConfig

	// nodes map[string]*Node

	// schemas map[string]*Schema

	// users    *userAuth
	listener net.Listener
	running  bool
}

func NewServer(cfg *config.ProxyConfig) (*Server, error) {
	s := new(Server)

	s.cfg = cfg

	var err error
	s.listener, err = net.Listen("tcp4", fmt.Sprintf(":%s", s.cfg.Global.Port))
	if err != nil {
		return nil, errors.Trace(err)
	}

	log.Notice("Dbatman Listen(tcp4) at [%s]", s.addr)
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
		if err := recover(); err != nil {
			const size = 4096
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			log.Fatal("onConn panic %v: %v\n%s", c.RemoteAddr().String(), err, buf)
		}

		session.Close()
	}()

	// Handshake error, here we do not need to close the conn
	if err := session.HandshakeWithFront(); err != nil {
		log.Warning("handshake error: %s", err.Error())
		return
	}

	session.front = mc

	session.Run()
}
