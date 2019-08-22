package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/markfaulk350/TrackPilotsAPI/entity"
	"github.com/markfaulk350/TrackPilotsAPI/service"
)

func UpdateUser(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Header().Set("Content-Type", "application/json")
		decoder := json.NewDecoder(r.Body)

		params := mux.Vars(r)
		userID := params["id"]

		// create pointer to empty user struct and fill with request body
		user := new(entity.User)
		if err := decoder.Decode(user); err != nil {
			fmt.Println("Unable to update user. Bad req body.")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(entity.JsonResponse{Success: false, Payload: "Unable to update user. Bad req body."})
			return
		}

		// send updated user struct to be inserted into database.
		err := svc.UpdateUser(userID, *user)
		if err != nil {
			fmt.Println("Could not update user info.")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.JsonResponse{Success: false, Payload: "Could not update user info."})
			return
		}

		// Create []byte from JsonResponse that returns success message
		jsonObj, err := json.Marshal(entity.JsonResponse{Success: true, Payload: ("Successfully updated user: " + userID)})
		if err != nil {
			fmt.Println("Failed marshalling json after updating user.")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonObj)
	}
}
