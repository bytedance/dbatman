package proxy

/*
import (
	"bytes"
	. "github.com/bytedance/dbatman/database/mysql"
)


func (session *Session) checkAuth(auth []byte) error {

	auths := session.server.getUserAuth(session.user)
	if auths == nil {
		return NewDefaultError(ER_ACCESS_DENIED_ERROR,
			session.conn.RemoteAddr().String(), session.user, "Yes")
	}

	for passwd, db := range auths.DB {
		if bytes.Equal(auth, CalcPassword(session.salt, []byte(passwd))) {
			// gotcha!!!
			session.db = db
			return nil
		}
	}
	return NewDefaultError(ER_ACCESS_DENIED_ERROR,
		session.conn.RemoteAddr().String(), session.user, "Yes")
}

func (session *Session) checkAuthWithDB(auth []byte, db string) error {
	var s *Schema
	if s = session.server.getSchema(db); s == nil {
		return NewDefaultError(ER_BAD_DB_ERROR, db)
	}

	if passwd, ok := s.auths[session.user]; !ok {
		return NewDefaultError(ER_ACCESS_DENIED_ERROR, session.conn.RemoteAddr().String(), session.user, "Yes")
	} else if !bytes.Equal(auth, CalcPassword(session.salt, []byte(passwd))) {
		return NewDefaultError(ER_ACCESS_DENIED_ERROR, session.conn.RemoteAddr().String(), session.user, "Yes")
	}

	if err := session.useDB(db); err != nil {
		return err
	}

	return nil
}
*/
