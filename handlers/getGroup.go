package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/markfaulk350/TrackPilotsAPI/entity"
	"github.com/markfaulk350/TrackPilotsAPI/service"
)

func GetGroup(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		groupID := params["id"]

		result, err := svc.GetGroup(groupID)
		if err != nil {
			fmt.Println("Get group failed")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.JsonResponse{Success: false, Payload: "Unable to get group data. Group might not exist."})
			return
		}

		jsonObj, err := json.Marshal(result)
		if err != nil {
			fmt.Println("Failed marshalling json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.JsonResponse{Success: false, Payload: "Unable to convert group data into JSON"})
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonObj)
		return
	}
}
