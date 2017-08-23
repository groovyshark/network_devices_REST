package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"strconv"
)

func getAllDevices(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(devices); err != nil {
		panic(err)
	}
}

func getSingleDevice(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)

	var deviceId int
	var err error
	if deviceId, err = strconv.Atoi(vars["deviceId"]); err != nil {
		panic(err)
	}

	device := findDevice(deviceId)
	if device.ID > 0 {

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(device); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)

	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

func indexPage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Welcome to Index page!")
}