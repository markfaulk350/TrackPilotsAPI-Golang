package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
	"github.com/markfaulk350/TrackPilotsAPI/service"
)

func GetAllUsers(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		result, err := svc.GetAllUsers()
		if err != nil {
			fmt.Println("Get all users failed")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.JsonResponse{Success: false, Payload: "Unable to get all users"})
			return
		}

		jsonObj, err := json.Marshal(result)
		if err != nil {
			fmt.Println("Failed marshalling []users json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.JsonResponse{Success: false, Payload: "Unable to convert all users data into JSON"})
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonObj)
		return
	}
}
