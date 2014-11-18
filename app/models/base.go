package models

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
)

type Base struct {
	
}

func (base *Base) InitDb() *gorp.DbMap {
	
	db, err := sql.Open("mysql", "root:daniel2912@/cmstest")
	
	if err != nil {
		panic("Error opening DB connection. Err")
		return nil
	}
	
	dbmap := &gorp.DbMap{Db:db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF-8"}}
	
	return dbmap
	
}
