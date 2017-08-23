package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		indexPage,
	},
	Route{
		"TodoIndex",
		"GET",
		"/devices",
		getAllDevices,
	},
	Route{
		"TodoShow",
		"GET",
		"/devices/{deviceId}",
		getSingleDevice,
	},
}
