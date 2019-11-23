package main

import (
	"log"
	"net/http"

	"github.com/RonaldCrb/go-mc/home"
	"github.com/RonaldCrb/go-mc/users"
	"github.com/julienschmidt/httprouter"
)

func main() {

	router := httprouter.New()
	router.GET("/", home.Main)
	router.GET("/healthCheck", home.HealthCheck)
	router.POST("/users", users.UserCreate)
	router.GET("/users", users.UserIndex)
	router.GET("/users/:id", users.UserFind)
	router.PUT("/users/:id", users.UserUpdate)
	router.DELETE("/users/:id", users.UserDelete)

	log.Fatal(http.ListenAndServe(":8080", router))
}
