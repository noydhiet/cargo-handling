package service

import (
	// "fmt"
	// "strings"
	"database/sql"
	// "log"
	// "net/http"
	// "text/template"
	_ "github.com/go-sql-driver/mysql"
)

// Connect to database
func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "PASSWORD"
	dbName := "db_go"
	dbIP := "192.168.20.9"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbIP+")/"+dbName)

	if err != nil {
		panic(err.Error())
	}

	return db
}


