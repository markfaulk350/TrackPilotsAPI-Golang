package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/markfaulk350/TrackPilotsAPI/utils"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
	"github.com/markfaulk350/TrackPilotsAPI/service"
	"github.com/rs/zerolog"
)

func CreateGroup(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
		decoder := json.NewDecoder(r.Body)

		group := new(entity.Group)
		if err := decoder.Decode(group); err != nil {
			msg := "Bad request body"
			logger.Debug().Err(err).Msg(msg)
			utils.RespondWithError(msg, err, http.StatusBadRequest, w)
			return
		}

		result, err := svc.CreateGroup(*group)
		if err != nil {
			msg := "Create group failed"
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
		w.Header().Set("Location", fmt.Sprintf("/groups/%v", result))
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonObj)
	}
}
