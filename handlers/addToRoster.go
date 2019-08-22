package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
	"github.com/markfaulk350/TrackPilotsAPI/service"
)

func AddToRoster(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Header().Set("Content-Type", "application/json")
		decoder := json.NewDecoder(r.Body)

		// create pointer to empty roster struct and fill with request body
		roster := new(entity.Roster)
		if err := decoder.Decode(roster); err != nil {
			fmt.Println("Unable to add user to group roster. Bad req body.")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(entity.JsonResponse{Success: false, Payload: "Unable to add user to group roster. Bad req body."})
			return
		}

		// send new roster struct to be inserted into database. Recieve new group ID as result
		err := svc.AddToRoster(*roster)
		if err != nil {
			fmt.Println("Could not add user to group in database.")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.JsonResponse{Success: false, Payload: "Could not add user to group in database."})
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(entity.JsonResponse{Success: true, Payload: "User added to group roster!"})
		return
	}
}
