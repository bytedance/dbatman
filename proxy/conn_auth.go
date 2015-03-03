package proxy

import (
	"bytes"
	. "github.com/wangjild/go-mysql-proxy/log"
	. "github.com/wangjild/go-mysql-proxy/mysql"
)

func (c *Conn) checkAuth(auth []byte) error {
	AppLog.Debug("checkAuth")
	auths := c.server.getUserAuth(c.user)
	if auths == nil {
		AppLog.Warn("connect without db, auths is nil")
		return NewDefaultError(ER_ACCESS_DENIED_ERROR, c.c.RemoteAddr().String(), c.user, "Yes")
	}

	for passwd, db := range auths.DB {
		if bytes.Equal(auth, CalcPassword(c.salt, []byte(passwd))) {
			// gotcha!!!
			c.db = db
			return nil
		}
	}
	return NewDefaultError(ER_ACCESS_DENIED_ERROR, c.c.RemoteAddr().String(), c.user, "Yes")
}

func (c *Conn) checkAuthWithDB(auth []byte, db string) error {
	var s *Schema
	if s = c.server.getSchema(db); s == nil {
		return NewDefaultError(ER_BAD_DB_ERROR, db)
	}

	if passwd, ok := s.auths[c.user]; !ok {
		return NewDefaultError(ER_ACCESS_DENIED_ERROR, c.c.RemoteAddr().String(), c.user, "Yes")
	} else if !bytes.Equal(auth, CalcPassword(c.salt, []byte(passwd))) {
		return NewDefaultError(ER_ACCESS_DENIED_ERROR, c.c.RemoteAddr().String(), c.user, "Yes")
	}

	if err := c.useDB(db); err != nil {
		return err
	}

	return nil
}
