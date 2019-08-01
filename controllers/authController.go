package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vasialek/VsLinks/data"
	"github.com/vasialek/VsLinks/helpers"
	"github.com/vasialek/VsLinks/models"
)

// AuthController to access it as class
type AuthController struct {
	userRepository *data.UserRepository
	authHelper     *helpers.AuthHelper
}

// NewAuthController returns instance of Auth controller
func NewAuthController() *AuthController {
	return &AuthController{
		userRepository: data.NewUserRepository(),
		authHelper:     helpers.NewAuthHelper(),
	}
}

// Login authenticates user POSTed form
func (ac *AuthController) Login(w http.ResponseWriter, rq *http.Request) {
	var loginRq models.LoginRequest
	if err := json.NewDecoder(rq.Body).Decode(&loginRq); err != nil {
		reportError(w, "Error decoding authentication request.", err)
		return
	}

	user, err := ac.userRepository.Login(loginRq.Email, loginRq.Password)
	if err != nil {
		reportError(w, "Could not log user in.", err)
		return
	}

	fmt.Println(user)
	jwt, err := ac.authHelper.GenerateJwt(&user, 10)
	if err != nil {
		reportError(w, "Could not generate JWT for authenicated user.", err)
		return
	}

	sendDataResponse(w, &models.LoginResponse{
		Status: true,
		Jwt:    jwt,
	})
}
