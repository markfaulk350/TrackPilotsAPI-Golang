package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
	"github.com/markfaulk350/TrackPilotsAPI/service"
)

func CreateGroup(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Header().Set("Content-Type", "application/json")
		decoder := json.NewDecoder(r.Body)

		// create pointer to empty group struct and fill with request body
		group := new(entity.Group)
		if err := decoder.Decode(group); err != nil {
			fmt.Println("Unable to create group. Bad req body.")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(entity.JsonResponse{Success: false, Payload: "Unable to create group. Bad req body."})
			return
		}

		// send new group struct to be inserted into database. Recieve new group ID as result
		result, err := svc.CreateGroup(*group)
		if err != nil {
			fmt.Println("Could not add group to database.")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.JsonResponse{Success: false, Payload: "Could not add group to database."})
			return
		}

		// Create []byte from JsonResponse that returns new group ID inside payload
		jsonObj, err := json.Marshal(entity.JsonResponse{Success: true, Payload: result})
		if err != nil {
			fmt.Println("Failed marshalling json after creating group.")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// return new group ID
		w.Header().Add("Location", fmt.Sprintf("/group/id/%v", result))
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonObj)
	}
}
