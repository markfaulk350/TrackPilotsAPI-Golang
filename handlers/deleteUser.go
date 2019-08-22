package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
	"github.com/markfaulk350/TrackPilotsAPI/service"
)

func DeleteUser(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		userID := params["id"]

		err := svc.DeleteUser(userID)
		if err != nil {
			fmt.Println("Unable to delete user:", userID)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.JsonResponse{Success: false, Payload: ("Unable to delete user: " + userID)})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(entity.JsonResponse{Success: true, Payload: ("User: " + userID + " has been deleted.")})
		return
	}
}
