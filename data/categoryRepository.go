package data

import (
	"github.com/aleksej.vasinov/visma-links/models"
	"github.com/zabawaba99/firego"
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
func (cr *CategoryRepository) GetAllActive() ([]models.LinkCategory, *error) {
	list := make([]models.LinkCategory, 1)
	list = append(list, models.LinkCategory{
		LinkCategoryID: "123-456-789",
		Name:           "Fake category",
	})
	return list, nil
}

func (cr *CategoryRepository) getDatabaseApp() (*firego.Firebase, *error) {
	return nil, nil
}
