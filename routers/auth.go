package routers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/vasialek/VsLinks/controllers"
	"github.com/vasialek/VsLinks/helpers"
	"github.com/vasialek/VsLinks/models"
)

// InitAuthRoutes initializes routes for authentication
func InitAuthRoutes(r *mux.Router) *mux.Router {
	ac := controllers.NewAuthController()

	r.HandleFunc("/api/v1/auth", ac.Login).Methods("POST")

	return r
}

func authenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, rq *http.Request) {
		fmt.Printf("Auth %s\n", rq.URL.RequestURI())
		if isAllowedAnonymously(rq.URL.RequestURI()) == false {
			authorizationHeader := rq.Header.Get("Authorization")
			if len(authorizationHeader) == 0 {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			ah := helpers.NewAuthHelper()
			u, err := ah.ValidateHeader(authorizationHeader)
			if err != nil {
				http.Error(w, "Invalid user", http.StatusUnauthorized)
				return
			}

			models.UserData = u
		}

		next.ServeHTTP(w, rq)
	})
}

func isAllowedAnonymously(url string) bool {
	if strings.HasSuffix(url, "/api/v1/auth") {
		return true
	}

	fmt.Printf("URL `%s` is not allowed to browse anonymously\n", url)
	return false
}
