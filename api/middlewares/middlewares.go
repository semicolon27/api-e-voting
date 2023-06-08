package middlewares

import (
	"errors"
	"log"
	"net/http"

	"github.com/semicolon27/api-e-voting/api/auth"
	"github.com/semicolon27/api-e-voting/api/responses"
)

func SetMiddlewareJSON(next http.HandlerFunc, title string) http.HandlerFunc {
	log.Println(title)
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAdminAuthentication(next http.HandlerFunc, title string) http.HandlerFunc {
	log.Println(title)
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenAdminValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized. 401"))
			return
		}
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc, title string) http.HandlerFunc {
	log.Println(title)
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenParticipantValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized. 401"))
			return
		}
		next(w, r)
	}
}
