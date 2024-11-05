package main

import (
	"testing"

	"github.com/rs/zerolog/log"
)

func TestGetSystemInfo(t *testing.T) {
	cmdr := NewCommander()
	info, err := cmdr.GetSystemInfo()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if info.Hostname == "" {
		t.Error("Expected hostname to be non-empty")
	}

	if info.IPAddress == "" {
		t.Error("Expected IP address to be non-empty")
	}

	log.Info().Msgf("result %v %s %s", info.Successful, info.Hostname, info.IPAddress)
}
