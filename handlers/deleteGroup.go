package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/markfaulk350/TrackPilotsAPI/entity"
	"github.com/markfaulk350/TrackPilotsAPI/service"
)

func DeleteGroup(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		groupID := params["id"]

		err := svc.DeleteGroup(groupID)
		if err != nil {
			fmt.Println("Unable to delete group:", groupID)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.JsonResponse{Success: false, Payload: ("Unable to delete group: " + groupID)})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(entity.JsonResponse{Success: true, Payload: ("Group: " + groupID + " has been deleted.")})
		return
	}
}
