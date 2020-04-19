package home

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Greeting ...
type Greeting struct {
	Message string
}

type CatchAll struct {
	Status    string
	TimeStamp string
	Origin    string
}

// HealthStatus ...
type HealthStatus struct {
	Status    string
	TimeStamp string
}

// Main => main endpoint
func Main(res http.ResponseWriter, req *http.Request) {
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
func HealthCheck(res http.ResponseWriter, req *http.Request) {
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

// CatchAllHandler => catch all endpoint for not found requests
func CatchAllHandler(res http.ResponseWriter, req *http.Request) {
	currentTime := time.Now()

	ca := CatchAll{
		Status:    "not found :(",
		TimeStamp: currentTime.Format("2006-01-02 15:04:05"),
		Origin:    req.RequestURI,
	}

	mca, err := json.Marshal(ca)
	if err != nil {
		fmt.Println(err)
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "%s\n", mca)
	fmt.Println("[HOME] => CatchAll function")
}
