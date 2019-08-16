package service

import (
	"log"
	"os"

	dt "cargo-handling/datastruct"
)

func UpdateStatusHandling(del dt.Handle) []dt.Handle {
	logger := log.NewLogfmtLogger(os.Stdout)
	logger.Log("Updating Database...")

	db := dbConn()
	upDB, err := db.Query("UPDATE ")

	if err != nil {
		panic(err.Error())
	}

	
}