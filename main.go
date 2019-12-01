package main

import (
	"log"
	"net/http"

	"github.com/RonaldCrb/go-pg-rest-api/home"
	"github.com/RonaldCrb/go-pg-rest-api/offers"
	"github.com/RonaldCrb/go-pg-rest-api/users"
	"github.com/julienschmidt/httprouter"
)

func main() {

	router := httprouter.New()

	// basic routes
	router.GET("/", home.Main)
	// healtcheck routes
	router.GET("/healthCheck", home.HealthCheck)
	// users routes
	router.POST("/users", users.UserCreate)
	router.GET("/users", users.UserIndex)
	router.GET("/users/:id", users.UserFind)
	router.PUT("/users/:id", users.UserUpdate)
	router.DELETE("/users/:id", users.UserDelete)
	// offers routes
	router.POST("/offers", offers.OfferCreate)
	router.GET("/offers", offers.OfferIndex)
	router.GET("/offers/:id", offers.OfferFind)
	router.PUT("/offers/:id", offers.OfferUpdate)
	router.DELETE("/offers/:id", offers.OfferDelete)

	log.Fatal(http.ListenAndServe(":8080", router))
}
