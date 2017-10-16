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

	encoder := json.NewEncoder(w)
	encoder.SetIndent("","\t")

	if err := encoder.Encode(devices); err != nil {
		panic(err)
	}
}

func getSingleDevice(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)

	var deviceId int
	var err error
	if deviceId, err = strconv.Atoi(vars["deviceId"]); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	device := findDevice(deviceId, devices)
	if device.ID > 0 {

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		encoder := json.NewEncoder(w)
		encoder.SetIndent("","\t")

		if err := encoder.Encode(device); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)

	encoder := json.NewEncoder(w)
	encoder.SetIndent("","\t")

	if err := encoder.Encode(jsonError{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func indexPage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Welcome to Index page!")
}