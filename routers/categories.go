package routers

import (
	"github.com/gorilla/mux"
	"github.com/vasialek/VsLinks/controllers"
)

// InitCategoryRoutes initializes routing for category CRUD
func InitCategoryRoutes(r *mux.Router) *mux.Router {
	cc := controllers.NewCategoriesController()

	r.HandleFunc("/api/v1/category", cc.CreateCategory).Methods("POST")
	r.HandleFunc("/api/v1/category", cc.GetActiveCategories).Methods("GET")

	return r
}
