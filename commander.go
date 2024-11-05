package main

import (
	"time"
)

type commander struct{}

func NewCommander() Commander {
	return &commander{}
}

type Commander interface {
	Ping(host string) (PingResult, error)
	GetSystemInfo() (SystemInfoResult, error)
}

type CommandRequest struct {
	Type    string `json:"type"`    // "ping" or "sysinfo"
	Payload string `json:"payload"` // For ping, this is the host
}

type CommandResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error,omitempty"`
}

type PingResult struct {
	Successful bool
	Time       time.Duration
}

type SystemInfoResult struct {
	Successful bool // I added, please see DOCS/Decisions.md item 3 for details
	Hostname   string
	IPAddress  string
}
