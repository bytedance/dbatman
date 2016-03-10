package sql

import (
	"testing"
)

func TestSet(t *testing.T) {
	var st IStatement

	st = testParse(`set global autocommit = 1`, t, false)
	matchType(t, st, &Set{})

	st = testParse(`set global autocommit = 1, sysvar = 2`, t, false)
	set := st.(*Set)

	v := set.VarList[0]
	if v.Life != Life_Global {
		t.Fatal("missed life type")
	}

	if v.Name != "autocommit" {
		t.Fatal("missed varname")
	}

	v = set.VarList[1]
	if v.Life != Life_Global {
		t.Fatal("missed life type")
	}

	if v.Name != "sysvar" {
		t.Fatal("missed varname")
	}

}

func TestShow(t *testing.T) {
	var st IStatement

	st = testParse(`show session variables like 'autocommit'`, t, false)
	matchType(t, st, &ShowVariables{})

	st = testParse(`show full tables in test`, t, false)
	matchSchemas(t, st, `test`)

	st = testParse(`show table status in test`, t, false)
	matchType(t, st, &ShowTableStatus{})
	matchSchemas(t, st, `test`)

	st = testParse(`show global status`, t, false)
	matchType(t, st, &ShowStatus{})

	st = testParse(`SHOW SLAVE STATUS`, t, false)
	matchType(t, st, &ShowSlaveStatus{})

	st = testParse(`SHOW SLAVE HOSTS`, t, false)
	matchType(t, st, &ShowSlaveHosts{})

	st = testParse(`SHOW Profiles`, t, false)
	matchType(t, st, &ShowProfiles{})

	st = testParse(`SHOW FULL PROCESSLIST`, t, false)
	matchType(t, st, &ShowProcessList{})

	st = testParse(`SHOW PLUGINS`, t, false)
	matchType(t, st, &ShowPlugins{})

	st = testParse(`SHOW PRIVILEGES`, t, false)
	matchType(t, st, &ShowPrivileges{})

	st = testParse(`SHOW OPEN TABLES IN test like 'tables_%'`, t, false)
	matchType(t, st, &ShowOpenTables{})
	matchSchemas(t, st, `test`)

	st = testParse(`SHOW MASTER STATUS`, t, false)
	matchType(t, st, &ShowMasterStatus{})

	st = testParse(`SHOW INDEX FROM mytable FROM mydb;`, t, false)
	matchType(t, st, &ShowIndex{})
	matchSchemas(t, st, `mydb`)

	st = testParse(`SHOW GRANTS FOR 'root'@'localhost';`, t, false)
	matchType(t, st, &ShowGrants{})

	st = testParse(`SHOW FUNCTION STATUS`, t, false)
	matchType(t, st, &ShowFunction{})

	st = testParse(`SHOW FUNCTION CODE dbname.func_name`, t, false)
	matchSchemas(t, st, `dbname`)

	st = testParse(`SHOW EVENTS FROM test;`, t, false)
	matchSchemas(t, st, `test`)

	st = testParse(`SHOW ERRORS`, t, false)
	matchType(t, st, &ShowErrors{})

	st = testParse(`SHOW COUNT(*) ERRORS`, t, false)
	matchType(t, st, &ShowErrors{})

	st = testParse(`Show STORAGE ENGINES`, t, false)
	matchType(t, st, &ShowEngines{})

	st = testParse(`SHOW ENGINE PERFORMANCE_SCHEMA STATUS`, t, false)
	matchType(t, st, &ShowEngines{})

	st = testParse(`SHOW Databases like '%presale%'`, t, false)
	matchType(t, st, &ShowDatabases{})

	st = testParse(`SHOW CREATE View test.view`, t, false)
	matchType(t, st, &ShowCreate{})
	matchSchemas(t, st, `test`)

	st = testParse(`SHOW CREATE TRIGGER test.trigger`, t, false)
	matchType(t, st, &ShowCreate{})
	matchSchemas(t, st, `test`)

	st = testParse(`SHOW CREATE TABLE test.table`, t, false)
	matchType(t, st, &ShowCreate{})
	matchSchemas(t, st, `test`)

	st = testParse(`SHOW CREATE EVENT test.e_daily`, t, false)
	matchType(t, st, &ShowCreate{})
	matchSchemas(t, st, `test`)

	st = testParse(`SHOW CREATE PROCEDURE test.simpleproc`, t, false)
	matchType(t, st, &ShowCreate{})
	matchSchemas(t, st, `test`)

	st = testParse(`SHOW CREATE DATABASE test`, t, false)
	matchType(t, st, &ShowCreateDatabase{})
	matchSchemas(t, st, `test`)

	st = testParse(`SHOW CHARACTER SET LIKE 'latin%';`, t, false)
	matchType(t, st, &ShowCharset{})

	st = testParse(`SHOW COLUMNS FROM mytable FROM mydb;`, t, false)
	matchSchemas(t, st, `mydb`)

	st = testParse(`SHOW COLUMNS FROM mydb.mytable;`, t, false)
	matchSchemas(t, st, `mydb`)

	st = testParse(`SHOW COLLATION LIKE 'latin1%';`, t, false)
	matchType(t, st, &ShowCollation{})

	st = testParse(`SHOW Binary LOGS;`, t, false)
	matchType(t, st, &ShowLogs{})

	st = testParse(`show binlog events in 'log1' from 123 limit 2, 4`, t, false)
	matchType(t, st, &ShowLogEvents{})
}

func TestTableMtStmt(t *testing.T) {
	st := testParse(`analyze table db1.tb1`, t, false)
	matchType(t, st, &Analyze{})
	matchSchemas(t, st, `db1`)

	st = testParse(`CHECK TABLE test.test_table FAST QUICK;`, t, false)
	matchType(t, st, &Check{})
	matchSchemas(t, st, `test`)

	st = testParse(`CHECKSUM TABLE test.test_table QUICK;`, t, false)
	matchType(t, st, &CheckSum{})
	matchSchemas(t, st, `test`)

	st = testParse(`OPTIMIZE TABLE foo.bar`, t, false)
	matchType(t, st, &Optimize{})
	matchSchemas(t, st, `foo`)

	st = testParse(`REPAIR NO_WRITE_TO_BINLOG  TABLE foo.bar quick`, t, false)
	matchType(t, st, &Repair{})
	matchSchemas(t, st, `foo`)

}

func TestPluginAndUdf(t *testing.T) {
	st := testParse(`CREATE AGGREGATE FUNCTION function_name RETURNS DECIMAL SONAME 'shared_library_name'`, t, false)
	matchType(t, st, &CreateUDF{})

	st = testParse(`INSTALL PLUGIN plugin_name SONAME 'shared_library_name'`, t, false)
	matchType(t, st, &Install{})
	if _, ok := st.(IPluginAndUdf); !ok {
		t.Fatalf("type[%T] is not a instance of IPluginAndUdf", st)
	}

	st = testParse(`UNINSTALL PLUGIN plugin_name`, t, false)
	matchType(t, st, &Uninstall{})
	if _, ok := st.(IPluginAndUdf); !ok {
		t.Fatalf("type[%T] is not a instance of IPluginAndUdf", st)
	}
}

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

func TestBinlog(t *testing.T) {
	st := testParse(`BINLOG 'str'`, t, false)
	matchType(t, st, &Binlog{})
}

func TestCacheIndex(t *testing.T) {
	st := testParse(`CACHE INDEX d1.t1, d2.t2, d3.t3 IN hot_cache;`, t, false)
	matchType(t, st, &CacheIndex{})
	matchSchemas(t, st, `d1`, `d2`, `d3`)

	st = testParse(`LOAD INDEX INTO CACHE pt PARTITION (p1, p3);`, t, false)
	matchType(t, st, &LoadIndex{})
	matchSchemas(t, st)

	st = testParse(`LOAD INDEX INTO CACHE db1.t1, db2.t2 IGNORE LEAVES;`, t, false)
	matchSchemas(t, st, `db1`, `db2`)
}

func TestFlush(t *testing.T) {
	st := testParse(`FLUSH TABLES db1.tbl_name , db2.tbl_name WITH READ LOCK`, t, false)
	matchType(t, st, &FlushTables{})
	matchSchemas(t, st, `db1`, `db2`)

	st = testParse(`flush logs`, t, false)
	matchType(t, st, &Flush{})
}

func TestKill(t *testing.T) {
	st := testParse(`kill connection 1234`, t, false)
	matchType(t, st, &Kill{})
}

func TestReset(t *testing.T) {
	st := testParse(`reset master, query cache, slave`, t, false)
	matchType(t, st, &Reset{})
}
