package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// GenericResponse returns timestamp and status to user in JSON
type GenericResponse struct {
	StatusCode int16
	Status     string
	TimeStamp  time.Time
}

// UserIndex => returns a slice of all users in the database
func UserIndex(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	usrs, err := AllUsers()

	if err != nil {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	uj, err := json.Marshal(usrs)
	if err != nil {
		log.Printf("[ERROR - USERS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "%s\n", uj)

}

// UserCreate creates a user instance in the database
func UserCreate(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	usr := User{}
	json.NewDecoder(req.Body).Decode(&usr)

	err := usr.CreateUser()
	if err != nil {
		log.Printf("[ERROR - USERS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	r := GenericResponse{
		StatusCode: 201,
		Status:     "User Created Succesfully",
		TimeStamp:  time.Now(),
	}

	rj, err := json.Marshal(r)
	if err != nil {
		log.Printf("[ERROR - USERS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	fmt.Fprintf(res, "%s\n", rj)
}

// UserFind finds a user instance by ID in the database
func UserFind(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(req)
	i := vars["id"]
	id, err := strconv.Atoi(i)
	if err != nil {
		log.Printf("[ERROR - USERS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(404), http.StatusNotFound)
		return
	}

	u := User{ID: id}

	usr, err := u.FindUser()

	if err != nil {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	uj, err := json.Marshal(usr)
	if err != nil {
		log.Printf("[ERROR - USERS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "%s\n", uj)
}

// UserUpdate Updates a user instance in the database
func UserUpdate(res http.ResponseWriter, req *http.Request) {
	if req.Method != "PUT" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(req)
	i := vars["id"]
	id, err := strconv.Atoi(i)
	if err != nil {
		log.Printf("[ERROR - USERS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(404), http.StatusNotFound)
		return
	}

	usr := User{ID: id}
	json.NewDecoder(req.Body).Decode(&usr)

	err = usr.UpdateUser()
	if err != nil {
		log.Printf("[ERROR - USERS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	r := GenericResponse{
		StatusCode: 202,
		Status:     "User Updated Succesfully",
		TimeStamp:  time.Now(),
	}

	rj, err := json.Marshal(r)
	if err != nil {
		log.Printf("[ERROR - USERS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(res, "%s\n", rj)
}

// UserDelete deletes a user instance from the database
func UserDelete(res http.ResponseWriter, req *http.Request) {
	if req.Method != "DELETE" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(req)
	i := vars["id"]
	id, err := strconv.Atoi(i)
	if err != nil {
		log.Printf("[ERROR - USERS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(404), http.StatusNotFound)
		return
	}

	u := User{ID: id}

	err = u.DeleteUser()

	if err != nil {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	r := GenericResponse{
		StatusCode: 201,
		Status:     "User succesfully DELETED",
		TimeStamp:  time.Now(),
	}

	rj, err := json.Marshal(r)
	if err != nil {
		log.Printf("[ERROR - USERS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	fmt.Fprintf(res, "%s\n", rj)
}
