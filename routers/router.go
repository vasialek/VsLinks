package routers

import "github.com/gorilla/mux"

// InitRoutes initializes routers for whole project
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	router = InitCategoryRoutes(router)
	router = InitLinkRoutes(router)

	return router
}
