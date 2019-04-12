package routers

import (
	"github.com/gorilla/mux"
	"github.com/vasialek/VsLinks/data/controllers"
)

// InitLinkRoutes initailizes routing to handle links
func InitLinkRoutes(r *mux.Router) *mux.Router {
	lr := controllers.NewLinksController()

	r.HandleFunc("/api/v1/link", lr.GetLinks)
	r.HandleFunc("/api/v1/link", lr.CreateLink)

	return r
}
