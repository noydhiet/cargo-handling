package main

import (
	"cargo-handling/transport"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
)

func main() {

	logger := log.NewLogfmtLogger(os.Stdout)

	transport.RegisterHttpServicesAndStartListener()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8989"
	}

	logger.Log("Listening-on", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logger.Log("listen.error", err)
	}
}
