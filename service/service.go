package service

import (
	"database/sql"

	dt "cargo-handling/datastruct"

	_ "github.com/go-sql-driver/mysql"
)

// Connect to database
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

func GetHandlingProcess(del dt.Handling) []dt.Handling {

	// logger := log.NewLogFmtLogger(os.Stdout)
	// logger.Log("Checking Database")

	db := dbConn()
	selDb, err := db.Query("SELECT * FROM t_trx_delivery WHERE routing_status = ?")

	if err != nil {
		panic(err.Error())
	}

	hndle := dt.Handling{}
	res := []dt.Handling{}

	for selDb.Next() {
		var routingStatus string
		err = selDb.Scan(&routingStatus)

		if err != nil {
			panic(err.Error())
		}

		hndle.ROUTING_STATUS = routingStatus
		res = append(res, hndle)
	}
	// defer db.Close()
	return res
}
