package main

import (
	"net"
	"os"
)

func (c *commander) GetSystemInfo() (SystemInfoResult, error) {
	hostname, err := os.Hostname()
	var systemInfoResult SystemInfoResult
	if err != nil {
		systemInfoResult.Successful = false
		return SystemInfoResult{}, err
	}

	addrs, err := net.LookupHost(hostname)
	if err != nil {
		systemInfoResult.Successful = false
		return SystemInfoResult{}, err
	}

	var ipAddresses string
	for _, ip := range addrs {
		ipAddresses = ipAddresses + ip + "\n"
	}
	systemInfoResult.Hostname = hostname
	systemInfoResult.IPAddress = ipAddresses
	return systemInfoResult, nil
}
