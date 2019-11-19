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

	usr := models.User{
		ID:        "1",
		FirstName: "Pedrito",
		LastName:  "Palotes, de los",
		Email:     "test@email.com",
	}

	uj, err := json.Marshal(usr)
	if err != nil {
		fmt.Println(err)
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "%s\n", uj)

}
