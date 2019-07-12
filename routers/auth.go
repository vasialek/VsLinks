package routers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// InitAuthRoutes initializes routes for authentication
func InitAuthRoutes(r *mux.Router) *mux.Router {
	r.HandleFunc("/api/v1/auth", func(w http.ResponseWriter, rq *http.Request) {
		w.Write([]byte("Login..."))
	}).Methods("POST")
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
