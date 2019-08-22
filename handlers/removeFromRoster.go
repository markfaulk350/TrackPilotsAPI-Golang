package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
	"github.com/markfaulk350/TrackPilotsAPI/service"
)

func RemoveFromRoster(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Header().Set("Content-Type", "application/json")
		decoder := json.NewDecoder(r.Body)

		// create pointer to empty roster struct and fill with request body
		roster := new(entity.Roster)
		if err := decoder.Decode(roster); err != nil {
			fmt.Println("Unable to remove user from group. Bad req body.")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(entity.JsonResponse{Success: false, Payload: "Unable to remove user from group. Bad req body."})
			return
		}

		// send roster struct to be removed from database. If no err then send success json response
		err := svc.RemoveFromRoster(*roster)
		if err != nil {
			fmt.Println("Could not remove user from group in database.")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.JsonResponse{Success: false, Payload: "Could not remove user from group in database."})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(entity.JsonResponse{Success: true, Payload: ("User has been removed from group.")})
		return
	}
}
