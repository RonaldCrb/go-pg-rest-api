package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RonaldCrb/go-mc/models"
	"github.com/julienschmidt/httprouter"
)

// UserController => exports the interface of a UserController
type UserController struct{}

// NewUserController => utility function to generate UserController pointer
func NewUserController() *UserController {
	return &UserController{}
}

// UserIndex => returns a slice of all users
func (uc UserController) UserIndex(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	if req.Method != "GET" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	usrs, err := models.AllUsers()

	if err != nil {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	uj, err := json.Marshal(usrs)
	if err != nil {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "%s\n", uj)

}
