package service

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "PASSWORD"
	dbName := "db_go"
	dbIp := "192.168.20.9"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbIp+")/"+dbName)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func viewALl(){

	db := dbConn()
	selDb, err := db.Query("SELECT ROUTING_STATUS, TRANS_STATUS FROM t_trx_delivery WHERE ROUTE_ID = ITENARY_ID")

	if err != nil{
		panic(err.Error())
	}
}


func Update(w http.ResponseWriter, r *http.Request){
	db := dbConn()

	updForm, err := db.Prepare("UPDATE t_trx_delivery SET ROUTING_STATUS = success, TRANSPORT_STATUS = success ")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}