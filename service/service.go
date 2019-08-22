package service

import (
	dt "cargo-handling/datastruct"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// connect to database
func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "db_go"
	dbIP := "127.0.0.1:3306"

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
		var fk_id_route_spec, fk_id_itenary, idt_delivery int
		var routing_status, transport_status, last_known_location string
		err = getHandle.Scan(&fk_id_route_spec, &fk_id_itenary, &idt_delivery, &routing_status, &transport_status, &last_known_location)

		if err != nil {
			panic(err.Error())
		}

		hndle.ROUTE_ID = fk_id_route_spec
		hndle.ITENARY_ID = fk_id_itenary
		hndle.DELIVERY_ID = idt_delivery
		hndle.ROUTING_STATUS = routing_status
		hndle.TRANSPORT_STATUS = transport_status
		hndle.LAST_KNOWN_LOCATION = last_known_location
		res = append(res, hndle)
	}

	return res

}
