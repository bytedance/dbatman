package proxy

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"io"

	. "github.com/bytedance/dbatman/database/mysql"
	"github.com/ngaut/log"
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

	var err error
	//check the global authip
	gc, err := session.config.GetGlobalConfig()
	cliAddr := session.cliAddr
	if gc.AuthIPActive == true {
		if len(gc.AuthIPs) > 0 {
			//TODO white and black ip logic
			globalAuthIp := &gc.AuthIPs
			authIpFlag := false
			for _, ip := range *globalAuthIp {
				if ip == cliAddr {
					authIpFlag = true
					break
				}
			}

			if authIpFlag != true {
				// log.Info("This user's Ip is not in the list of User's auth_Ip")
				return NewDefaultError(ER_NO, "IP Is not in the auth_ip list of the global config")
			}

		}
	}
	//global auth pass
	// There is no user named with parameter username
	if session.user, err = session.config.GetUserByName(username); err != nil {
		if session.user == nil {
			return NewDefaultError(ER_ACCESS_DENIED_ERROR, username, session.fc.RemoteAddr().String(), "Yes")
		}
		return NewDefaultError(ER_ACCESS_DENIED_ERROR, session.user.Username, session.fc.RemoteAddr().String(), "Yes")
	}

	if db != "" && session.user.DBName != db {
		log.Debugf("request db: %s, user's db: %s", db, session.user.DBName)
		return NewDefaultError(ER_BAD_DB_ERROR, db)
	}
	//TODO add the IP auth module to check the global auth_ip
	//check user config auth_IP with current Session Ip
	if gc.AuthIPActive == true {
		if len(session.user.AuthIPs) > 0 {
			userIPs := &session.user.AuthIPs
			authIpFlag := false
			log.Debug("client IP : ", session.cliAddr)
			log.Debug("User's Auth IP is: ", userIPs)
			for _, ip := range *userIPs {
				if cliAddr == ip {
					authIpFlag = true
					break
				}
			}
			if authIpFlag != true {
				log.Debug("This user's Ip is not in the list of User's auth_Ip")

				return NewDefaultError(ER_NO, "IP Is not in the auth_ip list of the user")
			}
		}
	}
	if !bytes.Equal(passwd, CalcPassword(session.salt, []byte(session.user.Password))) {
		return NewDefaultError(ER_ACCESS_DENIED_ERROR, session.user.Username, session.fc.RemoteAddr().String(), "Yes")
	}

	if err := session.useDB(session.user.DBName); err != nil {
		return err
	}

	return nil
}
