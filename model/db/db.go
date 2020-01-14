package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func NewDB(user, pwd, host, dbname string, port int) *sql.DB {
	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, pwd, host, port, dbname))
	if err != nil {
		panic(err.Error())
	}
	return conn
}
