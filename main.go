package main

import (
	"log"
	"net/http"

	"github.com/RonaldCrb/go-pg-rest-api/home"
	"github.com/RonaldCrb/go-pg-rest-api/offers"
	"github.com/RonaldCrb/go-pg-rest-api/users"
	"github.com/gorilla/mux"
)

func main() {
	// main router instantiation
	r := mux.NewRouter()

	// basic routes
	r.HandleFunc("/", home.Main).Methods("GET")
	// healtcheck routes
	r.HandleFunc("/healthCheck", home.HealthCheck).Methods("GET")
	// users routes
	r.HandleFunc("/users", users.UserCreate).Methods("POST")
	r.HandleFunc("/users", users.UserIndex).Methods("GET")
	r.HandleFunc("/users/{id}", users.UserFind).Methods("GET")
	r.HandleFunc("/users/{id}", users.UserUpdate).Methods("PUT")
	r.HandleFunc("/users/{id}", users.UserDelete).Methods("DELETE")
	// offers routes
	r.HandleFunc("/offers/createTable", offers.OfferCreateTable).Methods("GET")
	r.HandleFunc("/offers", offers.OfferCreate).Methods("POST")
	r.HandleFunc("/offers", offers.OfferIndex).Methods("GET")
	r.HandleFunc("/offers/{id}", offers.OfferFind).Methods("GET")
	r.HandleFunc("/offers/{id}", offers.OfferUpdate).Methods("PUT")
	r.HandleFunc("/offers/{id}", offers.OfferDelete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
