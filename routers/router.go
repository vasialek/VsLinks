package routers

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/gorilla/mux"
	"github.com/vasialek/VsLinks/models"
)

// InitRoutes initializes routers for whole project
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	router = InitCategoryRoutes(router)
	router = InitLinkRoutes(router)
	router = InitAuthRoutes(router)

	router.Use(authenticateMiddleware)

	if models.Settings.IsDevEnvironment {
		fmt.Println(dumpRoutes(router))
	}

	return router
}

func dumpRoutes(router *mux.Router) string {
	var buf bytes.Buffer
	router.Walk(func(route *mux.Route, r *mux.Router, children []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}

		methods, err := route.GetMethods()
		if err != nil {
			// Display that no HTTP method is assigned
			methods = []string{"**ANY**"}
		}

		buf.WriteString(fmt.Sprintf("%-20s", strings.Join(methods, ", ")))
		buf.WriteString(t + "\n")

		return nil
	})

	return buf.String()
}
