package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// InitCategoryRoutes initializes routing for category CRUD
func InitCategoryRoutes(r *mux.Router) *mux.Router {
	r.HandleFunc("/api/v1/category", func(w http.ResponseWriter, rq *http.Request) {
		w.Write([]byte("Category list..."))
	})

	return r
}
