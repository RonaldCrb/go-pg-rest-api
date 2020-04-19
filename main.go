package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/RonaldCrb/go-pg-rest-api/home"
	"github.com/RonaldCrb/go-pg-rest-api/users"
	"github.com/gorilla/handlers"
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

	fmt.Println("Cerberus ready to GO!")
	// implement logging middleware, initiate server in port, log fatal in case of error
	log.Fatal(http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, r)))

}
