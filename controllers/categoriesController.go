package controllers

import (
	"net/http"

	"github.com/vasialek/VsLinks/data"
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
