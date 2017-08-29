package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
)

func testGetAllDevices(t *testing.T) {
	req, err := http.NewRequest("GET", "/devices", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(getAllDevices)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	testDevices := localDevices()
	expected, _ := json.MarshalIndent(testDevices, "","\t")

	// Check the response body is what we expect.
	if rr.Body.String() != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func testSingleDevice(t *testing.T) {
	req, err := http.NewRequest("GET", "/devices/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(getAllDevices)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	testDevices := localDevices()
	expected, _ := json.MarshalIndent(testDevices[0], "","\t")

	// Check the response body is what we expect.
	if rr.Body.String() != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

