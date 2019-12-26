package sqlutil

type TestDB struct {
	DriverName string
	ConnStr    string
}

var TestDB_Sqlite3 = TestDB{DriverName: "sqlite3", ConnStr: ":memory:"}
var TestDB_Mysql = TestDB{DriverName: "mysql", ConnStr: "smartems:password@tcp(localhost:3306)/smartems_tests?collation=utf8mb4_unicode_ci"}
var TestDB_Postgres = TestDB{DriverName: "postgres", ConnStr: "user=smartemstest password=smartemstest host=localhost port=5432 dbname=smartemstest sslmode=disable"}
var TestDB_Mssql = TestDB{DriverName: "mssql", ConnStr: "server=localhost;port=1433;database=smartemstest;user id=smartems;password=Password!"}
