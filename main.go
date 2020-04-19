package main

import (
	"log"
	"net/http"

	"github.com/RonaldCrb/go-pg-rest-api/home"
	"github.com/RonaldCrb/go-pg-rest-api/users"
	"github.com/gorilla/mux"
)

func main() {
	// Here we are instantiating the gorilla/mux router
	r := mux.NewRouter()

	// On the default page we will simply serve our static index page.
	r.Handle("/", http.FileServer(http.Dir("./views/")))

	// We will setup our server so we can serve static assest like images, css from the /static/{file} route
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// healtcheck routes
	r.HandleFunc("/healthCheck", home.HealthCheck).Methods("GET")

	// users routes
	r.HandleFunc("/users", users.UserCreate).Methods("POST")
	r.HandleFunc("/users", users.UserIndex).Methods("GET")
	r.HandleFunc("/users/{id}", users.UserFind).Methods("GET")
	r.HandleFunc("/users/{id}", users.UserUpdate).Methods("PUT")
	r.HandleFunc("/users/{id}", users.UserDelete).Methods("DELETE")
	r.HandleFunc("/auth/register", users.Register).Methods("POST")
	r.HandleFunc("/auth/login", users.Login).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", r))
}
