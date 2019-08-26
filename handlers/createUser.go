package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
	"github.com/markfaulk350/TrackPilotsAPI/service"
	"github.com/markfaulk350/TrackPilotsAPI/utils"
	"github.com/rs/zerolog"
)

func CreateUser(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
		decoder := json.NewDecoder(r.Body)

		user := new(entity.User)
		if err := decoder.Decode(user); err != nil {
			msg := "Bad request body"
			logger.Debug().Err(err).Msg(msg)
			utils.RespondWithError(msg, err, http.StatusBadRequest, w)
			return
		}

		result, err := svc.CreateUser(*user)
		if err != nil {
			msg := "Adding user to database failed"
			logger.Debug().Err(err).Msg(msg)
			utils.RespondWithError(msg, err, http.StatusInternalServerError, w)
			return
		}

		utils.RespondWithSuccess("User has been created", result, http.StatusCreated, w)
		return
	}
}
