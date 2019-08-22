package service

import (
	"database/sql"

	dt "cargo-handling/datastruct"
)

// connect to database
func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "db_go"
	dbIP := "127.0.0.1"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbIP+")/"+dbName)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func GetHandlingProcess(get dt.Handle) []dt.Handle {

	db := dbConn()
	getHandle, err := db.Query("SELECT * FROM t_trx_delivery WHERE routing_status = ?", get.ROUTING_STATUS)

	if err != nil {
		panic(err.Error())
	}

	hndle := dt.Handle{}
	res := []dt.Handle{}

	for getHandle.Next() {
		var routingStatus string
		err = getHandle.Scan(&routingStatus)

		if err != nil {
			panic(err.Error())
		}

		hndle.ROUTING_STATUS = routingStatus
		res = append(res, hndle)
	}
	return res
}
