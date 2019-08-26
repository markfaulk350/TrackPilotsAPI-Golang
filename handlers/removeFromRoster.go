package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
	"github.com/markfaulk350/TrackPilotsAPI/service"
	"github.com/markfaulk350/TrackPilotsAPI/utils"
	"github.com/rs/zerolog"
)

func RemoveFromRoster(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
		decoder := json.NewDecoder(r.Body)

		roster := new(entity.Roster)
		if err := decoder.Decode(roster); err != nil {
			msg := "Bad request body"
			logger.Debug().Err(err).Msg(msg)
			utils.RespondWithError(msg, err, http.StatusBadRequest, w)
			return
		}

		if err := svc.RemoveFromRoster(*roster); err != nil {
			switch err {
			case sql.ErrNoRows:
				msg := "Failed to remove user from group. user or group does not exist"
				logger.Debug().Err(err).Msg(msg)
				utils.RespondWithError(msg, err, http.StatusInternalServerError, w)
				return
			default:
				msg := "Failed to remove user from group"
				logger.Debug().Err(err).Msg(msg)
				utils.RespondWithError(msg, err, http.StatusInternalServerError, w)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	}
}
