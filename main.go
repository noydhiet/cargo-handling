package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Server started on: http://localhost:8989/")
	http.ListenAndServe(":8989", nil)
	// GetHandlingProcess()
}
