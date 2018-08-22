package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dsbell81/go-pgsql-api/datastore"
	"github.com/dsbell81/go-pgsql-api/models"
	"github.com/dsbell81/go-pgsql-api/utils"
)

type (

	//For Post - /users/register
	UserResource struct {
		Data models.User `json:"data"`
	}

	//For Post - /users/login
	LoginResource struct {
		Data LoginModel `json:"data"`
	}

	//Response for authorized account Post - /accounts/login
	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}

	//For authentication
	LoginModel struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	//For authorized user with access token
	AuthUserModel struct {
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}
)

// Register add a new user
// Handler for HTTP Post - "/users/register"
func Register(w http.ResponseWriter, r *http.Request) {
	var newUserResource UserResource
	// Decode the incoming User json
	err := json.NewDecoder(r.Body).Decode(&newUserResource)
	if err != nil {
		utils.DisplayAppError(w, err,
			"Invalid User data",
			500,
		)
		return
	}

	newUserData := &newUserResource.Data

	//check password for minimum requirements
	if newUserData.Password == "" {
		err = errors.New("Password is blank")
	}

	if err != nil {
		utils.DisplayAppError(w, err,
			"Password does not meet requirements",
			500,
		)
		return
	}

	// Insert user record
	err = datastore.CreateUser(newUserData)
	if err != nil {
		utils.DisplayAppError(w, err,
			"An unexpected datastore error has occurred",
			500,
		)
		return
	}

	//clear password from response
	newUserData.Password = ""

	j, err := json.Marshal(UserResource{Data: *newUserData})
	if err != nil {
		utils.DisplayAppError(w, err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

// Login authenticates the HTTP request with username and apssword
// Handler for HTTP Post - "/users/login"
func Login(w http.ResponseWriter, r *http.Request) {
	//coming soon
}
