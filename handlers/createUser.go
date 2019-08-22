package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
	"github.com/markfaulk350/TrackPilotsAPI/service"
)

func CreateUser(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Header().Set("Content-Type", "application/json")
		decoder := json.NewDecoder(r.Body)

		// create pointer to empty user struct and fill with request body
		user := new(entity.User)
		if err := decoder.Decode(user); err != nil {
			fmt.Println("Unable to create user. Bad req body.")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(entity.JsonResponse{Success: false, Payload: "Unable to create user. Bad req body."})
			return
		}

		// send new user struct to be inserted into database. Recieve new user ID as result
		result, err := svc.CreateUser(*user)
		if err != nil {
			fmt.Println("Could not add user to database.")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.JsonResponse{Success: false, Payload: "Could not add user to database."})
			return
		}

		// Create []byte from JsonResponse that returns new user ID inside payload
		jsonObj, err := json.Marshal(entity.JsonResponse{Success: true, Payload: result})
		if err != nil {
			fmt.Println("Failed marshalling json after creating user.")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// return new user ID
		w.Header().Add("Location", fmt.Sprintf("/user/id/%v", result))
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonObj)
	}
}
