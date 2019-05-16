package routers

import (
	"github.com/gorilla/mux"
	"github.com/vasialek/VsLinks/controllers"
)

// InitLinkRoutes initailizes routing to handle links
func InitLinkRoutes(r *mux.Router) *mux.Router {
	lr := controllers.NewLinksController()

	r.HandleFunc("/api/v1/links", lr.GetLinks)
	r.HandleFunc("/api/v1/links", lr.CreateLink).Methods("POST")
	r.HandleFunc("/api/v1/links/{linkid}/category/{categoryid}", lr.SetLinkCategory).Methods("POST")

	return r
}
