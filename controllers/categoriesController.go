package controllers

import (
	"net/http"
	"strings"

	"github.com/vasialek/VsLinks/data"
	"github.com/vasialek/VsLinks/helpers"
	"github.com/vasialek/VsLinks/models"
)

// CategoriesController used as class
type CategoriesController struct {
	repository *data.CategoryRepository
}

// NewCategoriesController returns instance of controller
func NewCategoriesController() *CategoriesController {
	return &CategoriesController{
		repository: data.NewCategoryRepository(),
	}
}

// GetActiveCategories returns JSON with list of active categories
func (cc *CategoriesController) GetActiveCategories(w http.ResponseWriter, rq *http.Request) {
	categories, err := cc.repository.GetAllActive()
	if err != nil {
		reportError(w, "Error getting list of categories", err)
		return
	}

	sendDataResponse(w, categories)
}

// CreateCategory creates Link category or returns errors
func (cc *CategoriesController) CreateCategory(w http.ResponseWriter, rq *http.Request) {
	category := &models.LinkCategory{}

	if err := helpers.Decode(rq, &category); err != nil {
		reportError(w, "Error decoding Link category to create", err)
		return
	}

	errors := cc.validate(category)
	if len(errors) > 0 {
		reportError(w, strings.Join(errors, ". "), nil)
		return
	}

	category, err := cc.repository.CreateCategory(*category)
	if err != nil {
		reportError(w, "Error creating Link category", err)
		return
	}

	sendDataResponse(w, category)
}

func (cc *CategoriesController) validate(category *models.LinkCategory) (errors []string) {
	// var errors []string

	if len(category.Name) < 2 {
		errors = append(errors, "Link category name should be at least 2 symbols")
	} else if len(category.Name) > 256 {
		errors = append(errors, "Link category name should be less than 256 symbols")
	}

	return errors
}
