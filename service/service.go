package service

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	// "fmt"
	"github.com/go-kit/kit/log"
	"os"
	// "strings"

	dt "cargo-handling/datastruct"
)

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


func UpdateStatusHandling(del dt.Handle) []dt.Handle {
	logger := log.NewLogfmtLogger(os.Stdout)
	logger.Log("Updating Database...")

	db := dbConn()

	upDB, err := db.Query(`UPDATE t_trx_delivery AS a INNER JOIN t_mtr_route_spec AS b on a.fk_id_route_spec = b.id_route_spec SET a.routing_status = "Success" WHERE a.last_known_location = b.destination`)						

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// upDB.Exec("Success","Success")

	hand := dt.Handle{}
	res := []dt.Handle{}

	for upDB.Next() {
		var id_route_spec, id_itenary int
		var routing_status, transport_status, destination, last_known_location string
		err = upDB.Scan(&routing_status, &transport_status, &id_itenary, &id_route_spec, &destination, &last_known_location)
		if err != nil {
			panic(err.Error())
		}
		// hand.ID_ITENARY = id_itenary
		hand.ROUTING_STATUS = routing_status
		hand.TRANSPORT_STATUS = transport_status
		// hand.ID_ROUTE = id_route_spec
		res = append(res, hand)
	}

	return res
}