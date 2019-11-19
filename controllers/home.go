package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// HomeController ...
type HomeController struct{}

// Greeting ...
type Greeting struct {
	Message string
}

// HealthCheck ...
type HealthCheck struct {
	Message   string
	TimeStamp string
}

// NewHomeController => utility function to generate HomeController pointer
func NewHomeController() *HomeController {
	return &HomeController{}
}

// Home => main endpoint
func (hc HomeController) Home(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	g := Greeting{Message: "Hello Visitor"}

	gj, err := json.Marshal(g)
	if err != nil {
		fmt.Println(err)
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "%s\n", gj)
}

// HealthCheck => utility endpoint for healthchecks
func (hc HomeController) HealthCheck(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	currentTime := time.Now()

	g := HealthCheck{
		Message:   "Hello Visitor",
		TimeStamp: currentTime.Format("2006-01-02 15:04:05"),
	}

	gj, err := json.Marshal(g)
	if err != nil {
		fmt.Println(err)
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "%s\n", gj)
}
