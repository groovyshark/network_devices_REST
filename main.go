package main

import (
	"log"
	"net/http"
)

var devices Devices

func handleRequests() {
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	devices = localDevices()
	handleRequests()
}
