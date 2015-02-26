package sql

import (
	"testing"
)

func TestAccountMgrStmt(t *testing.T) {
	st := testParse(`ALTER USER 'jeffrey'@'localhost' PASSWORD EXPIRE;`, t, false)
	matchType(t, st, &AlterUser{})

	st = testParse(`CREATE USER 'jeffrey'@'localhost' IDENTIFIED BY 'mypass';`, t, false)
	matchType(t, st, &CreateUser{})

	st = testParse(`DROP USER 'jeffrey'@'localhost';`, t, false)
	matchType(t, st, &DropUser{})

	st = testParse(`GRANT SELECT ON db2.invoice TO 'jeffrey'@'localhost';`, t, false)
	matchType(t, st, &Grant{})

	st = testParse(`RENAME USER 'jeffrey'@'localhost' TO 'jeff'@'127.0.0.1';`, t, false)
	matchType(t, st, &RenameUser{})

	st = testParse(`REVOKE INSERT ON *.* FROM 'jeffrey'@'localhost';`, t, false)
	matchType(t, st, &Revoke{})

	st = testParse(`SET PASSWORD FOR 'jeffrey'@'localhost' = PASSWORD('cleartext password');`, t, false)
	// matchType(t, st, &SetPassword{})
}
