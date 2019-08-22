package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/caarlos0/env"
	"github.com/gorilla/mux"
	"github.com/markfaulk350/TrackPilotsAPI/dbclient"
	"github.com/markfaulk350/TrackPilotsAPI/handlers"
	"github.com/markfaulk350/TrackPilotsAPI/service"
)

func Start() {

	CONN_PORT := os.Getenv("CONN_PORT")

	config := dbclient.Config{}
	env.Parse(&config)
	svc := service.New(&config)

	r := mux.NewRouter()
	s := r.PathPrefix("/trackingAPI/v1").Subrouter()

	// Health Check
	s.HandleFunc("/health", handlers.Health()).Methods(http.MethodGet)
	// Users
	s.HandleFunc("/users", handlers.GetAllUsers(svc)).Methods(http.MethodGet)
	s.HandleFunc("/users", handlers.CreateUser(svc)).Methods(http.MethodPost)
	s.HandleFunc("/users/{id}", handlers.GetUser(svc)).Methods(http.MethodGet)
	s.HandleFunc("/users/{id}", handlers.UpdateUser(svc)).Methods(http.MethodPut)
	s.HandleFunc("/users/{id}", handlers.DeleteUser(svc)).Methods(http.MethodDelete)
	// Groups
	s.HandleFunc("/groups", handlers.GetAllGroups(svc)).Methods(http.MethodGet)
	s.HandleFunc("/groups", handlers.CreateGroup(svc)).Methods(http.MethodPost)
	s.HandleFunc("/groups/{id}", handlers.GetGroup(svc)).Methods(http.MethodGet)
	s.HandleFunc("/groups/{id}", handlers.UpdateGroup(svc)).Methods(http.MethodPut)
	s.HandleFunc("/groups/{id}", handlers.DeleteGroup(svc)).Methods(http.MethodDelete)
	// Group Roster
	s.HandleFunc("/roster", handlers.AddToRoster(svc)).Methods(http.MethodPost)
	s.HandleFunc("/roster", handlers.RemoveFromRoster(svc)).Methods(http.MethodDelete)
	s.HandleFunc("/roster/{id}", handlers.GetGroupRoster(svc)).Methods(http.MethodGet)
	// Tracking Data
	s.HandleFunc("/grouptrackingdata/{id}", handlers.GetGroupTrackingData(svc)).Methods(http.MethodGet)

	fmt.Println("Server listening on port", CONN_PORT)
	log.Fatal(http.ListenAndServe(":"+CONN_PORT, r))

}

// sudo nano .bash_profile
// export ENVNAME="envVar"
