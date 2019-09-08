package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/markfaulk350/TrackPilotsAPI/entity"
	"github.com/markfaulk350/TrackPilotsAPI/service"
)

func GetGroupTrackingData(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Eventually I need to take in parameters to determine how much tracking data I want to send back for a given time period
		// For now Im just returning all tracking data

		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		groupID := params["id"]

		result, err := svc.GetGroupTrackingData(groupID)
		if err != nil {
			fmt.Println("Get groups users pings failed.")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.JsonResponse{Success: false, Payload: "Unable to get groups users pings."})
			return
		}

		jsonObj, err := json.Marshal(result)
		if err != nil {
			fmt.Println("Failed marshalling json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.JsonResponse{Success: false, Payload: "Unable to convert groups location data into JSON"})
			return
		}

		//fmt.Println("jsoned data is ", string(jsonObj))

		w.WriteHeader(http.StatusOK)
		w.Write(jsonObj)
		return
	}
}
