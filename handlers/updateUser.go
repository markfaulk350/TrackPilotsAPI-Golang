package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/markfaulk350/TrackPilotsAPI/entity"
	"github.com/markfaulk350/TrackPilotsAPI/service"
	"github.com/markfaulk350/TrackPilotsAPI/utils"
	"github.com/rs/zerolog"
)

func UpdateUser(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
		decoder := json.NewDecoder(r.Body)

		params := mux.Vars(r)
		userID := params["id"]

		user := new(entity.User)
		if err := decoder.Decode(user); err != nil {
			msg := "Bad request body"
			logger.Debug().Err(err).Msg(msg)
			utils.RespondWithError(msg, err, http.StatusBadRequest, w)
			return
		}

		err := svc.UpdateUser(userID, *user)
		if err != nil {
			msg := "Update user failed"
			logger.Error().Err(err).Msg(msg)
			utils.RespondWithError(msg, err, http.StatusInternalServerError, w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	}
}
