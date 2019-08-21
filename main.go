package main

import (
	"cargo-handling/transport"
	"log"
	"net/http"
	"os"
)

func main() {

	logger := log.NewLogfmtLogger(os.Stdout)

	transport.RegisterHttpServicesAndStartListener()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8989"
	}
	logger.log("Listening-on", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logger.log("listen.error", err)
	}
}
