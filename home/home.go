package home

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Greeting ...
type Greeting struct {
	Message string
}

// HealthStatus ...
type HealthStatus struct {
	Status    string
	TimeStamp string
}

// Main => main endpoint
func Main(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	g := Greeting{Message: "Hello Visitor"}

	gj, err := json.Marshal(g)
	if err != nil {
		fmt.Println(err)
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "%s\n", gj)
	fmt.Println("[HOME] => Main function")
}

// HealthCheck => utility endpoint for healthchecks
func HealthCheck(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	currentTime := time.Now()

	g := HealthStatus{
		Status:    "GO-MC Healthy!",
		TimeStamp: currentTime.Format("2006-01-02 15:04:05"),
	}

	gj, err := json.Marshal(g)
	if err != nil {
		fmt.Println(err)
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "%s\n", gj)
	fmt.Println("[HOME] => HealthCheck function")

}
