package handlers

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"

	"github.com/markfaulk350/TrackPilotsAPI/service"
	"github.com/markfaulk350/TrackPilotsAPI/utils"
)

func DeleteUser(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
		params := mux.Vars(r)
		userID := params["id"]

		if err := svc.DeleteUser(userID); err != nil {
			switch err {
			case sql.ErrNoRows:
				msg := "Delete user failed. User not found"
				logger.Error().Err(err).Msg(msg)
				utils.RespondWithError(msg, err, http.StatusNotFound, w)
				return
			default:
				msg := "Delete user failed"
				logger.Error().Err(err).Msg(msg)
				utils.RespondWithError(msg, err, http.StatusInternalServerError, w)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	}
}
