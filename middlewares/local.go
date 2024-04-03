package middlewares

import (
	"bloglist/helpers"
	"context"
	"net/http"
	"strings"
)

func UserExtractor(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		tokenString = tokenString[len("Bearer "):]
		user, err := helpers.ParseToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		//TODO: maybe check if user exists in database

		ctxWithUser := context.WithValue(r.Context(), "User", user)
		r = r.WithContext(ctxWithUser)

		h(w, r)
	}
}
