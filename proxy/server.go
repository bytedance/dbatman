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
	cfg *config.Config

	nodes map[string]*Node

	schemas map[string]*Schema

	users    *userAuth
	listener net.Listener
	running  bool
}

func NewServer(cfg *config.Config) (*Server, error) {
	s := new(Server)

	s.cfg = cfg

	s.addr = cfg.Addr
	s.user = cfg.User
	s.password = cfg.Password

	if err := s.parseNodes(); err != nil {
		return nil, errors.Trace(err)
	}

	if err := s.parseSchemas(); err != nil {
		return nil, errors.Trace(err)
	}

	if err := s.parseUserAuths(); err != nil {
		return nil, errors.Trace(err)
	}

	var err error
	s.listener, err = net.Listen("tcp4", s.addr)
	if err != nil {
		return nil, errors.Trace(err)
	}

	log.Notice("Dbatman Listen(tcp4) at [%s]", s.addr)
	return s, nil
}

func (s *Server) Run() error {
	s.running = true

	for s.running {
		conn, err := s.listener.Accept()
		if err != nil {
			SysLog.Warn("accept error %s", err.Error())
			continue
		}

		go s.onConn(conn)
	}

	return nil
}

func (s *Server) Close() {
	s.running = false
	if s.listener != nil {
		s.listener.Close()
	}
}

func (s *Server) onConn(c net.Conn) {
	conn := s.newConn(c)

	defer func() {
		if err := recover(); err != nil {
			const size = 4096
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			AppLog.Warn("onConn panic %v: %v\n%s", c.RemoteAddr().String(), err, buf)
		}

		conn.Close()
	}()

	if err := conn.handshake(); err != nil {
		AppLog.Warn("handshake error %s", err.Error())
		c.Close()
		return
	}

	conn.Run()
}
