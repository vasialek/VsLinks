package models

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// UserData stores information about logged in user
var UserData *User

// TypeOfResource use to set status and type of resources, like LinkCategory
type TypeOfResource int

const (
	// ActivePublic active and available to all users
	ActivePublic TypeOfResource = iota + 1
	// ActivePrivate active, but only for owner
	ActivePrivate
	// DeletedPublic deleted
	DeletedPublic
	// DeletedPrivate deleted
	DeletedPrivate
)

// #region JSON requests/responses

// Response is base JSON response
type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// LoginRequest model to authenticate user
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	ClientID string `json:"client_id"`
}

// LoginResponse model for successful or not login
type LoginResponse struct {
	Status bool   `json:"status"`
	Jwt    string `json:"jwt"`
}

// #endregion

// User represents user who is able to use VsLinks system
type User struct {
	UserID   string         `json:"user_id"`
	StatusID TypeOfResource `json:"status_id"`
	Name     string         `json:"name"`
	Email    string         `json:"email"`
	Password string         `json:"password"`
}

// UserClaims stores user attributes
type UserClaims struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

// Link is model for create/transfer links
type Link struct {
	LinkID         string    `json:"link_id"`
	UserID         string    `json:"user_id"`
	TypeID         int       `json:"type_id"`
	LinkCategoryID string    `json:"link_category_id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	URL            string    `json:"url"`
	Rate           int       `json:"rate"`
	CreatedAt      time.Time `json:"created_at"`
}

// LinkCategory to store information about categories
type LinkCategory struct {
	LinkCategoryID string         `json:"link_category_id"`
	StatusID       TypeOfResource `json:"status_id"`
	UserID         string         `json:"user_id"`
	Name           string         `json:"name"`
	CreatedAt      time.Time      `json:"created_at"`
}
