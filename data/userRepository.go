package data

import (
	"errors"

	"github.com/vasialek/VsLinks/helpers"
	"github.com/vasialek/VsLinks/models"
	"github.com/zabawaba99/firego"
)

// UserRepository provides access to User table as object
type UserRepository struct {
	db *firego.Firebase
}

// NewUserRepository returns instance of UserRepository
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// GetAllUsers returns list of all users from database
func (ur *UserRepository) GetAllUsers() ([]models.User, error) {
	// TODO: create database
	users := make([]models.User, 2)
	users[0] = models.User{
		UserID:   "2677b0d2-009b-414b-92da-f8d5cc65efa1",
		StatusID: models.ActivePublic,
		Name:     "Aleksej V.",
		Email:    "proglamer@gmail.com",
		Password: "$2a$10$ooKHvSTAPVoAmZZceJxgv.pPsK9lPSPcyigjpy/2l/G8w1Way194K",
	}
	users[1] = models.User{
		UserID:   "7c79c65f-57d4-4e2b-be6b-7591ede15f0a",
		StatusID: models.ActivePublic,
		Name:     "Andrej N.",
		Email:    "anikolskij@gmail.com",
		Password: "$2a$10$EvXnbkZ.uFZcuk6Jnxg1i.cpRCmLA1gILdRcDXIVY2H6265K7K05u",
	}

	return users, nil
}

// Login returns logged in user or error
func (ur *UserRepository) Login(email, password string) (user models.User, err error) {
	// Till w/o database
	users, _ := ur.GetAllUsers()

	for _, u := range users {
		if u.Email == email && helpers.CheckPasswordHash(u.Password, password) {
			return u, nil
		}
	}

	return user, errors.New("Incorrect email/password")
}
