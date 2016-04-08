package proxy

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	. "github.com/bytedance/dbatman/database/mysql"
	"github.com/ngaut/log"
	"io"
)

func RandomBuf(size int) ([]byte, error) {

	buf := make([]byte, size)
	if debug {
		for i, _ := range buf {
			buf[i] = 0x01
		}

		return buf, nil
	}

	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return nil, err
	}

	for i, b := range buf {
		if uint8(b) == 0 {
			buf[i] = '0'
		}
	}
	return buf, nil
}

func CalcPassword(scramble, password []byte) []byte {
	if len(password) == 0 {
		return nil
	}

	// stage1Hash = SHA1(password)
	crypt := sha1.New()
	crypt.Write(password)
	stage1 := crypt.Sum(nil)

	// scrambleHash = SHA1(scramble + SHA1(stage1Hash))
	// inner Hash
	crypt.Reset()
	crypt.Write(stage1)
	hash := crypt.Sum(nil)

	// outer Hash
	crypt.Reset()
	crypt.Write(scramble)
	crypt.Write(hash)
	scramble = crypt.Sum(nil)

	// token = scrambleHash XOR stage1Hash
	for i := range scramble {
		scramble[i] ^= stage1[i]
	}
	return scramble
}

func (session *Session) CheckAuth(username string, passwd []byte, db string) error {

	log.Debugf("user login: name=%s db=%s", username, db)

	var err error

	// There is no user named with parameter username
	if session.user, err = session.config.GetUserByName(username); err != nil {
		return NewDefaultError(ER_ACCESS_DENIED_ERROR, session.user.Username, session.fc.RemoteAddr().String(), "Yes")
	}

	if db != "" && session.user.DBName != db {
		return NewDefaultError(ER_BAD_DB_ERROR, db)
	}

	if !bytes.Equal(passwd, CalcPassword(session.salt, []byte(session.user.Password))) {
		return NewDefaultError(ER_ACCESS_DENIED_ERROR, session.user.Username, session.fc.RemoteAddr().String(), "Yes")
	}

	if err := session.useDB(session.user.DBName); err != nil {
		return err
	}

	return nil
}
