package data

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/vasialek/VsLinks/models"
	"github.com/zabawaba99/firego"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// CategoryRepository provides access to category table
type CategoryRepository struct {
	db *firego.Firebase
}

// NewCategoryRepository returns instance of CategoryRepository
func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{}
}

// GetAllActive returns list of categories for links
func (cr *CategoryRepository) GetAllActive() ([]models.LinkCategory, error) {
	fb, err := cr.getDatabaseApp()
	if err != nil {
		log.Printf("CategoryRepository::GetAllActive: Error connecting to Firebase. %s\n", err)
		return nil, err
	}

	val := map[string]models.LinkCategory{}
	if err = fb.Child("category").Value(&val); err != nil {
		log.Printf("CategoryRepository::GetAllActive: %s\n", err)
		return nil, errors.New("Error getting list of active categories")
	}

	log.Printf("got %d active categories\n", len(val))
	categories := make([]models.LinkCategory, len(val))
	pos := 0

	for _, v := range val {
		categories[pos] = v
		pos++
	}

	return categories, nil
}

func (cr *CategoryRepository) getDatabaseApp() (*firego.Firebase, error) {
	if cr.db == nil {
		fmt.Println("Connecting to Firebase...")
		d, err := ioutil.ReadFile("./keys/vsm-links-db-firebase-adminsdk-gubgg-98be377e48.json")
		if err != nil {
			log.Printf("getDatabaseApp: can't read Firebase key: %s\n", err)
			return nil, err
		}

		conf, err := google.JWTConfigFromJSON(d, "https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/firebase.database")
		if err != nil {
			return nil, err
		}

		cr.db = firego.New("https://vsm-links-db.firebaseio.com/", conf.Client(oauth2.NoContext))
	}
	return cr.db, nil
}
