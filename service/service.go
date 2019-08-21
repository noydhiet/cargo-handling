package service

import (
	"database/sql"

	dt "cargo-handling/datastruct"
)

// connect to database
func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := ""
	dbPass := ""
	dbName := "db_go"
	dbIP := "127.0.0.1"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbIP+")/"+dbName)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func GetHandleProcess(get dt.Handling) []dt.Handling {

	db := dbConn()
	getHandle, err := db.Query("SELECT * FROM t_trx_delivery WHERE routing_status = ?")

	if err != nil {
		panic(err.Error())
	}

	hndle := dt.Handling{}
	res := []dt.Handling{}

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
