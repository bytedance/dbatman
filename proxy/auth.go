package proxy

import (
	"fmt"
)

type userAuth struct {
	Map map[string]*passDB
}

func newUserAuth() *userAuth {
	s := new(userAuth)
	s.Map = make(map[string]*passDB)
	return s
}

type passDB struct {
	DB map[string]string
}

func newPassDB() *passDB {
	pass := new(passDB)
	pass.DB = make(map[string]string)
	return pass
}

func (s *passDB) add(user string, passwd string, db string) error {
	if _, ok := s.DB[passwd]; ok {
		return fmt.Errorf("user[%s] with same passwd has multi db, this is forbidden!")
	}

	s.DB[passwd] = db
	return nil
}

func (s *Server) getUserAuth(user string) *passDB {
	return s.users.Map[user]
}

func (s *Server) parseUserAuths() error {
	uas := newUserAuth()
	for _, schemaCfg := range s.cfg.Schemas {
		for _, auth := range schemaCfg.Auths {
			if _, ok := uas.Map[auth.User]; !ok {
				uas.Map[auth.User] = newPassDB()
			}

			if err := uas.Map[auth.User].add(auth.User, auth.Passwd, schemaCfg.DB); err != nil {
				return err
			}
		}
	}

	s.users = uas

	return nil
}
