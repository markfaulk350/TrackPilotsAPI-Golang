package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/markfaulk350/TrackPilotsAPI/entity"
)

func Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(entity.JsonResponse{Success: true, Payload: "Tracking API is up and running!"})
		return
	}
}
