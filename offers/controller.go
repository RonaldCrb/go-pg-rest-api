package offers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// GenericResponse type of responses for POST => /offers, PUT => /offers/{id}, DELETE => /offers/{id}
type GenericResponse struct {
	StatusCode int16
	Status     string
	TimeStamp  time.Time
}

// IndexResponse type for responses to GET => /offers
type IndexResponse struct {
	StatusCode int16
	Status     string
	TimeStamp  time.Time
	Data       []Offer
	DataSize   int
}

// FindResponse type of response for GET => /offers/{id}
type FindResponse struct {
	StatusCode int16
	Status     string
	TimeStamp  time.Time
	Data       Offer
	DataSize   int
}

// OfferIndex => returns a slice of all Offers in the database
func OfferIndex(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	ofrs, err := AllOffers()

	if err != nil {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	resp := IndexResponse{
		StatusCode: 200,
		Status:     "List of offers found ([]Offer)",
		TimeStamp:  time.Now(),
		DataSize:   len(ofrs),
		Data:       ofrs,
	}

	response, err := json.Marshal(resp)

	if err != nil {
		log.Printf("[ERROR - OFFERS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "%s\n", response)

}

// OfferCreate creates a Offer instance in the database
func OfferCreate(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	ofr := Offer{}
	json.NewDecoder(req.Body).Decode(&ofr)

	err := ofr.CreateOffer()
	if err != nil {
		log.Printf("[ERROR - OFFERS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	r := GenericResponse{
		StatusCode: 201,
		Status:     "Offer Created Succesfully",
		TimeStamp:  time.Now(),
	}

	rj, err := json.Marshal(r)
	if err != nil {
		log.Printf("[ERROR - OFFERS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	fmt.Fprintf(res, "%s\n", rj)
}

// OfferFind finds a Offer instance by ID in the database
func OfferFind(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(req)
	i := vars["id"]
	id, err := strconv.Atoi(i)
	if err != nil {
		log.Printf("[ERROR - OFFERS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(404), http.StatusNotFound)
		return
	}

	o := Offer{ID: id}

	ofr, err := o.FindOffer()

	if err != nil {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	resp := FindResponse{
		StatusCode: 201,
		Status:     "Offer found (type Offer)",
		TimeStamp:  time.Now(),
		Data:       ofr,
		DataSize:   1,
	}

	response, err := json.Marshal(resp)
	if err != nil {
		log.Printf("[ERROR - OFFERS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "%s\n", response)
}

// OfferUpdate Updates a Offer instance in the database
func OfferUpdate(res http.ResponseWriter, req *http.Request) {
	if req.Method != "PUT" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(req)
	i := vars["id"]
	id, err := strconv.Atoi(i)
	if err != nil {
		log.Printf("[ERROR - OFFERS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(404), http.StatusNotFound)
		return
	}

	ofr := Offer{ID: id}
	json.NewDecoder(req.Body).Decode(&ofr)

	err = ofr.UpdateOffer()
	if err != nil {
		log.Printf("[ERROR - OFFERS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	r := GenericResponse{
		StatusCode: 202,
		Status:     "Offer Updated Succesfully",
		TimeStamp:  time.Now(),
	}

	rj, err := json.Marshal(r)
	if err != nil {
		log.Printf("[ERROR - OFFERS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(res, "%s\n", rj)
}

// OfferDelete deletes a Offer instance from the database
func OfferDelete(res http.ResponseWriter, req *http.Request) {
	if req.Method != "DELETE" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(req)
	i := vars["id"]
	id, err := strconv.Atoi(i)
	if err != nil {
		log.Printf("[ERROR - OFFERS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(404), http.StatusNotFound)
		return
	}

	o := Offer{ID: id}

	err = o.DeleteOffer()

	if err != nil {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	r := GenericResponse{
		StatusCode: 201,
		Status:     "Offer succesfully DELETED",
		TimeStamp:  time.Now(),
	}

	rj, err := json.Marshal(r)
	if err != nil {
		log.Printf("[ERROR - OFFERS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	fmt.Fprintf(res, "%s\n", rj)
}
