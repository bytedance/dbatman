package proxy

import (
	"fmt"
	"github.com/wangjild/go-mysql-proxy/config"
)

type Schema struct {
	db    string
	node  *Node
	auths map[string]string
}

func (s *Server) parseSchemas() error {
	s.schemas = make(map[string]*Schema)

	for _, schemaCfg := range s.cfg.Schemas {
		if _, ok := s.schemas[schemaCfg.DB]; ok {
			return fmt.Errorf("duplicate schema [%s].", schemaCfg.DB)
		}

		n := s.getNode(schemaCfg.Node)
		if n == nil {
			return fmt.Errorf("schema [%s] node [%s] config is not exists.", schemaCfg.DB, schemaCfg.Node)
		}

		auths, err := s.getAuths(schemaCfg)
		if err != nil {
			return err
		}

		s.schemas[schemaCfg.DB] = &Schema{
			db:    schemaCfg.DB,
			node:  n,
			auths: auths,
		}
	}

	return nil
}

func (s *Server) getAuths(schema config.SchemaConfig) (map[string]string, error) {
	if len(schema.Auths) == 0 {
		return nil, fmt.Errorf("schema [%s]'s auth is empty.", schema.DB)
	}

	auth := make(map[string]string)

	for _, v := range schema.Auths {
		if _, ok := auth[v.User]; ok {
			return nil, fmt.Errorf("schema [%s] has duplicate user[%s]", schema.DB, v.User)
		}

		auth[v.User] = v.Passwd
	}

	return auth, nil
}

func (s *Server) getSchema(db string) *Schema {
	return s.schemas[db]
}
