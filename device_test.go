package main

import "testing"

func testFindDevice(t *testing.T) {

	testDevices := localDevices()
	device := findDevice(1, testDevices)

	if device.ID != 1 {
		t.Errorf("Expected 1, got %v", device.ID)
	}
}