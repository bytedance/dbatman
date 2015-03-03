package proxy

import (
	"github.com/wangjild/go-mysql-proxy/config"
	. "github.com/wangjild/go-mysql-proxy/log"
	"net"
	//	"runtime"
)

type Server struct {
	cfg *config.Config

	addr     string
	user     string
	password string

	running bool

	listener net.Listener

	nodes map[string]*Node

	schemas map[string]*Schema

	users *userAuth
}

func NewServer(cfg *config.Config) (*Server, error) {
	s := new(Server)

	s.cfg = cfg

	s.addr = cfg.Addr
	s.user = cfg.User
	s.password = cfg.Password

	if err := s.parseNodes(); err != nil {
		return nil, err
	}

	if err := s.parseSchemas(); err != nil {
		return nil, err
	}

	if err := s.parseUserAuths(); err != nil {
		return nil, err
	}

	var err error
	s.listener, err = net.Listen("tcp4", s.addr)
	if err != nil {
		return nil, err
	}

	SysLog.Notice("Go-MySQL-Proxy Listen(tcp4) at [%s]", s.addr)
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
		/*if err := recover(); err != nil {
			const size = 4096
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			AppLog.Warn("onConn panic %v: %v\n%s", c.RemoteAddr().String(), err, buf)
		}*/

		conn.Close()
	}()

	if err := conn.Handshake(); err != nil {
		AppLog.Warn("handshake error %s", err.Error())
		c.Close()
		return
	}

	conn.Run()
}
