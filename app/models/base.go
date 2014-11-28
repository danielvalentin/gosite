package models

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/coopernurse/gorp"
)

func getDB() *gorp.DbMap {
	conn, err := sql.Open("mysql", "root:daniel2912@/godb")
	if err != nil {
		panic("cant connect to db")
	}
	
	dbmap := &gorp.DbMap{Db: conn, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	
	dbmap.AddTableWithName(User{}, "users").SetKeys(true, "Id")
	
	return dbmap
}
