package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/markfaulk350/TrackPilotsAPI/entity"
	"github.com/markfaulk350/TrackPilotsAPI/service"
)

func UpdateGroup(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Header().Set("Content-Type", "application/json")
		decoder := json.NewDecoder(r.Body)

		params := mux.Vars(r)
		groupID := params["id"]

		// create pointer to empty group struct and fill with request body
		group := new(entity.Group)
		if err := decoder.Decode(group); err != nil {
			fmt.Println("Unable to update group. Bad req body.")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(entity.JsonResponse{Success: false, Payload: "Unable to update group. Bad req body."})
			return
		}

		// send updated group struct to be inserted into database.
		err := svc.UpdateGroup(groupID, *group)
		if err != nil {
			fmt.Println("Could not update group info.")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.JsonResponse{Success: false, Payload: "Could not update group info."})
			return
		}

		// Create []byte from JsonResponse that returns success message
		jsonObj, err := json.Marshal(entity.JsonResponse{Success: true, Payload: ("Successfully updated group: " + groupID)})
		if err != nil {
			fmt.Println("Failed marshalling json after updating group.")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonObj)
	}
}
