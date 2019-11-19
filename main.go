package main

import (
	"log"
	"net/http"

	"github.com/RonaldCrb/go-mc/controllers"
	"github.com/julienschmidt/httprouter"
)

// database connection

func main() {
	homeController := controllers.NewHomeController()
	userController := controllers.NewUserController()

	router := httprouter.New()
	router.GET("/", homeController.Home)
	router.GET("/healthCheck", homeController.HealthCheck)
	router.GET("/users", userController.UserIndex)

	log.Fatal(http.ListenAndServe(":8080", router))
}
