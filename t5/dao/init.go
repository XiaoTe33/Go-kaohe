package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Db = InitMysql()

func InitMysql() *sql.DB {
	dsn := "root:root@tcp(localhost:3306)/kh"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil
	}
	err = db.Ping()
	if err != nil {
		return nil
	}
	return db
}
