package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)
func Sqlclient() *sql.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "12345678"
	dbName := "sms"
	// database connection configuration
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=True" )
	if err != nil {
		log.Fatal("error connecting DB : ", err.Error())
	}
	return db
}
