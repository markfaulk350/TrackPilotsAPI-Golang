package app

import (
	"log"
	"net/http"

	"github.com/apex/gateway"
	"github.com/caarlos0/env"

	// gmuxHandlers "github.com/gorilla/handlers"
	gmuxHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/markfaulk350/TrackPilotsAPI/dbclient"
	"github.com/markfaulk350/TrackPilotsAPI/handlers"
	"github.com/markfaulk350/TrackPilotsAPI/service"
)

func Start() {

	// Comment for Prod
	// CONN_PORT := os.Getenv("CONN_PORT")

	config := dbclient.Config{}
	env.Parse(&config)
	svc := service.New(&config)

	allowedHeaders := gmuxHandlers.AllowedHeaders([]string{"X-Requested-With"})
	allowedOrigins := gmuxHandlers.AllowedOrigins([]string{"*"})
	allowedMethods := gmuxHandlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

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

	// Comment for Prod
	// fmt.Println("Server listening on port", CONN_PORT)
	// log.Fatal(http.ListenAndServe(":"+CONN_PORT, r))

	// Uncomment for Prod
	// log.Fatal(gateway.ListenAndServe("", r))
	// or with gzip
	// log.Fatal(gateway.ListenAndServe("", gmuxHandlers.CompressHandler(r)))
	// or with CORS
	// log.Fatal(gateway.ListenAndServe("", gmuxHandlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(r)))
	// or both
	log.Fatal(gateway.ListenAndServe("", gmuxHandlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(gmuxHandlers.CompressHandler(r))))

	// With CORS
	// log.Fatal(http.ListenAndServe(":"+CONN_PORT, gmuxHandlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(r)))

	// With gzip
	// log.Fatal(http.ListenAndServe(":"+CONN_PORT, gmuxHandlers.CompressHandler(r)))

}

// sudo nano .bash_profile
// export ENVNAME="envVar"
