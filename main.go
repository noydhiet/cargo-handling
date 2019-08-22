package main

import (
	_ "expvar"
	"os"
	"github.com/go-kit/kit/log"
	"net/http"
	"cargo-handling/transport"
)

func main() {
	// log.Println("Server started on: http://localhost:8989/")
	// http.ListenAndServe(":8989", nil)

	logger := log.NewLogfmtLogger(os.Stdout)
	transport.RegisterHttpsServicesAndStartListener()
	//port := os.Getenv("PORT")
	//if port == "" {
	//	port = "8080"
	//}
	port := "8989"
	logger.Log("listening-on", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logger.Log("listen.error", err)
		//fmt.Println("Error")
	}
}