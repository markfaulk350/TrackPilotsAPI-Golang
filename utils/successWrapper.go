package utils

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/rs/zerolog"
)

type SuccessWrapper struct {
	Status  string `json:"status"`
	Msg     string `json:"message"`
	Payload string `json:"payload"`
}

// Facilitates handlers in the sending of success responses over http
func RespondWithSuccess(msg string, payload string, statusCode int, w http.ResponseWriter) {
	jsonObj, err := json.Marshal(SuccessWrapper{Status: http.StatusText(statusCode), Msg: msg, Payload: payload})
	if err != nil {
		msg := "Unable to send successful response"
		logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
		logger.Error().Err(err).Msg(msg)
		RespondWithError(msg, err, http.StatusInternalServerError, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonObj)
}
