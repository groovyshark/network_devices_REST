package main

import (
	"net"
	"fmt"
	"log"
)

type Device struct {
	ID 			 int 	`json:"ID"`
	Name 		 string `json:"Name"`
	HardwareAddr string `json:"HardwareAddr"`
	NetworkAddr  string `json:"NetworkAddr"`
}

type Devices []Device


func findDevice(id int) Device {
	for _, device := range devices {
		if device.ID == id {
			return device
		}
	}
	// Return empty Device if not found
	return Device{}
}


func localDevices() Devices {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Print(fmt.Errorf("Interfaces: %v\n", err.Error()))
		return Devices{}
	}

	id := 1
	var devices Devices
	for _, i := range interfaces {

		addrs, err := i.Addrs()
		if err != nil {
			log.Print(fmt.Errorf("Addrs: %v\n", err.Error()))
			continue
		}

		name := i.Name

		var hardAddr string
		if hardAddr = i.HardwareAddr.String(); hardAddr == "" {
			hardAddr = "none"
		}

		var strAddrs string
		for _, a := range addrs {
			strAddrs += fmt.Sprintf("%v:%v ", a.Network(), a.String())
		}

		devices = append(devices, Device{ ID: id,
										  Name: name,
										  HardwareAddr: hardAddr,
										  NetworkAddr: strAddrs })
		id += 1
	}
	return devices
}
