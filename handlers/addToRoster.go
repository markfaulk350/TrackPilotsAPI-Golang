package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
	"github.com/markfaulk350/TrackPilotsAPI/service"
	"github.com/markfaulk350/TrackPilotsAPI/utils"
	"github.com/rs/zerolog"
)

func AddToRoster(svc service.Service) http.HandlerFunc {
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

		result, err := svc.AddToRoster(*roster)
		if err != nil {
			msg := "Failed to add user to group roster"
			logger.Debug().Err(err).Msg(msg)
			utils.RespondWithError(msg, err, http.StatusInternalServerError, w)
			return
		}

		jsonObj, err := json.Marshal(entity.JsonResponse{Success: true, Payload: result})
		if err != nil {
			msg := "Failed marshaling json"
			logger.Debug().Err(err).Msg(msg)
			utils.RespondWithError(msg, err, http.StatusInternalServerError, w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Location", fmt.Sprintf("/roster/%v", result))
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonObj)
		return
	}
}
