package proxy

import (
	"fmt"
	"github.com/bytedance/dbatman/config"
	_ "github.com/bytedance/dbatman/database/mysql"
	"github.com/juju/errors"
	"github.com/ngaut/log"
	"net"
	"runtime"
)

// Server is the proxy server. It handle the request from frontend, process and dispatch
// queries, picking right backend conn due to the request context.
type Server struct {
	cfg *config.Conf

	// nodes map[string]*Node

	// schemas map[string]*Schema

	// users    *userAuth
	listener net.Listener
	running  bool
}

func NewServer(cfg *config.Conf) (*Server, error) {
	s := new(Server)

	s.cfg = cfg

	var err error

	port := s.cfg.GetConfig().Global.Port
	s.listener, err = net.Listen("tcp4", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, errors.Trace(err)
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
		log.Warnf("handshake error: %s", err.Error())
		return
	}

	log.Debugf("handshake successful!")

	if err := session.Run(); err != nil {
		// TODO
		// session.WriteError(NewDefaultError(err))
	}
}
