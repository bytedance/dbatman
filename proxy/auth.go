package proxy

import (
	"bytes"
	"github.com/bytedance/dbatman/config"
	. "github.com/bytedance/dbatman/database/mysql"
)

func (session *Session) CheckAuth(username string, passwd []byte, db string) error {
	var user *config.UserConfig

	// There is no user named with parameter username
	if user = session.config.GetUserByName(username); user != nil {
		return NewDefaultError(ER_ACCESS_DENIED_ERROR, session.fc.RemoteAddr().String(), session.user, "Yes")
	}

	if db != "" && user.DBName != db {
		return NewDefaultError(ER_BAD_DB_ERROR, db)
	}

	if !bytes.Equal(passwd, CalcPassword(session.salt, []byte(passwd))) {
		return NewDefaultError(ER_ACCESS_DENIED_ERROR, session.fc.RemoteAddr().String(), session.user, "Yes")
	}

	if err := session.useDB(db); err != nil {
		return err
	}

	return nil
}
