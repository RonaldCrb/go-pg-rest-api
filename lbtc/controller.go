package lbtc

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

// GenericResponse returns timestamp and status to Offer in JSON
type GenericResponse struct {
	StatusCode int16
	Status     string
	TimeStamp  time.Time
}

// OfferIndex => returns a slice of all Offers in the database
func OfferIndex(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	if req.Method != "GET" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	ofrs, err := AllOffers()

	if err != nil {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	oj, err := json.Marshal(ofrs)
	if err != nil {
		log.Fatal(err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "%s\n", oj)

}

// OfferCreate creates a Offer instance in the database
func OfferCreate(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	if req.Method != "POST" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	ofr := Offer{}
	json.NewDecoder(req.Body).Decode(&ofr)

	err := ofr.CreateOffer()
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	fmt.Fprintf(res, "%s\n", rj)
}

// OfferFind finds a Offer instance by ID in the database
func OfferFind(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	if req.Method != "GET" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	i := params.ByName("id")
	id, err := strconv.Atoi(i)
	if err != nil {
		log.Fatal(err)
		http.Error(res, http.StatusText(404), http.StatusNotFound)
		return
	}

	o := Offer{ID: id}

	ofr, err := o.FindOffer()

	if err != nil {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	oj, err := json.Marshal(ofr)
	if err != nil {
		log.Fatal(err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "%s\n", oj)
}

// OfferUpdate Updates a Offer instance in the database
func OfferUpdate(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	if req.Method != "PUT" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	i := params.ByName("id")
	id, err := strconv.Atoi(i)
	if err != nil {
		log.Fatal(err)
		http.Error(res, http.StatusText(404), http.StatusNotFound)
		return
	}

	ofr := Offer{ID: id}
	json.NewDecoder(req.Body).Decode(&ofr)

	err = ofr.UpdateOffer()
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(res, "%s\n", rj)
}

// OfferDelete deletes a Offer instance from the database
func OfferDelete(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	if req.Method != "DELETE" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	i := params.ByName("id")
	id, err := strconv.Atoi(i)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	fmt.Fprintf(res, "%s\n", rj)
}
