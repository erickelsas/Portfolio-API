package middlewares

import (
	"log"
	"net/http"
	"portfolio-api/src/authorization"
	"portfolio-api/src/responses"
)

func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunc(w, r)
	}
}

func Auth(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := authorization.TokenValidate(r); erro != nil {
			responses.Error(w, http.StatusUnauthorized, erro)
			return
		}

		nextFunc(w, r)
	}
}
