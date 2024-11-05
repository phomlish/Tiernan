package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

func main() {

	log.Debug().Msg("tiernan started")

	commander := NewCommander()
	server := &http.Server{
		Addr:    ":8081",
		Handler: handleRequests(commander),
	}
	log.Fatal().Msgf("%v", server.ListenAndServe())
}

func handleRequests(cmdr Commander) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/execute", handleCommand(cmdr))
	return mux
}

func handleCommand(cmdr Commander) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var commandResponse CommandResponse
		var commandRequest CommandRequest
		if r.Method != "POST" {
			errorMessage := fmt.Sprintf("only POSTs are supported, your request was a %s", r.Method)
			commandResponse.Error = errorMessage
			commandResponse.Success = false
			err := json.NewEncoder(w).Encode(commandResponse)
			if err != nil {
				log.Info().Msgf("Error sending response %v", err)
			}
		}

		err := json.NewDecoder(r.Body).Decode(&commandRequest)
		if err != nil {
			commandResponse.Error = "unable to decode request"
			commandResponse.Success = false
			err = json.NewEncoder(w).Encode(commandResponse)
			if err != nil {
				log.Info().Msgf("Error sending response %v", err)
			}
			return
		}

		log.Info().Msgf("handling %s", commandRequest.Type)
		// a bit overkill to use a switch here, I predict the bosses will want more request types
		switch commandRequest.Type {
		case "ping":
			pingResult, err := cmdr.Ping(commandRequest.Payload)
			if err != nil {
				commandResponse.Error = err.Error()
				log.Info().Msgf("Error processing request %v", err)
			}
			commandResponse.Data = pingResult
		case "sysinfo":
			systemInfoResult, err := cmdr.GetSystemInfo()
			if err != nil {
				commandResponse.Error = err.Error()
				log.Info().Msgf("Error processing request %v", err)
			}
			commandResponse.Data = systemInfoResult
		default:
			errorMessage := fmt.Sprintf("didn't recognize request %s", commandRequest.Type)
			http.Error(w, errorMessage, http.StatusBadRequest)
			return
		}

		commandResponse.Success = true

		err = json.NewEncoder(w).Encode(commandResponse)
		if err != nil {
			log.Info().Msgf("Error sending response %v", err)
		}

	}
}
