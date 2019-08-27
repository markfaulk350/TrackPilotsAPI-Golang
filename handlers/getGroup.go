package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/markfaulk350/TrackPilotsAPI/service"
	"github.com/markfaulk350/TrackPilotsAPI/utils"
	"github.com/rs/zerolog"
)

func GetGroup(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
		params := mux.Vars(r)
		groupID := params["id"]

		result, err := svc.GetGroup(groupID)
		if err != nil {
			switch err.(type) {
			case service.ProfileNotFoundError:
				msg := "Group not found"
				logger.Error().Err(err).Msg(msg)
				utils.RespondWithError(msg, err, http.StatusNotFound, w)
				return
			default:
				msg := "Get group failed"
				logger.Error().Err(err).Msg(msg)
				utils.RespondWithError(msg, err, http.StatusInternalServerError, w)
			}
		}

		jsonObj, err := json.Marshal(result)
		if err != nil {
			msg := "Failed marshaling json"
			logger.Error().Err(err).Msg(msg)
			utils.RespondWithError(msg, err, http.StatusInternalServerError, w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonObj)
		return
	}
}
