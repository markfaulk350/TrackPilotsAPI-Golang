package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/markfaulk350/TrackPilotsAPI/entity"
	"github.com/markfaulk350/TrackPilotsAPI/service"
)

func GetUser(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		userID := params["id"]

		result, err := svc.GetUser(userID)
		if err != nil {
			fmt.Println("Get user failed")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.JsonResponse{Success: false, Payload: "Unable to get user data. User might not exist."})
			return
		}

		jsonObj, err := json.Marshal(result)
		if err != nil {
			fmt.Println("Failed marshalling json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.JsonResponse{Success: false, Payload: "Unable to convert user data into JSON"})
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonObj)
		return
	}
}
