package main

import (
	"testing"

	"github.com/rs/zerolog/log"
)

func TestPing(t *testing.T) {
	cmdr := NewCommander()

	info, err := cmdr.Ping("google.com")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	log.Info().Msgf("result %v %s", info.Successful, info.Time)
}
