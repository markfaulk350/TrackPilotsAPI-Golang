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

func GetGroupRoster(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
		params := mux.Vars(r)
		groupID := params["id"]

		result, err := svc.GetGroupRoster(groupID)
		if err != nil {
			msg := "Get group roster failed"
			logger.Error().Err(err).Msg(msg)
			utils.RespondWithError(msg, err, http.StatusInternalServerError, w)
			return
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
