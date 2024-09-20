package utils

import (
	"errors"
	"net/http"
	"strings"
)

func getJwtToken(r *http.Request) string {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return ""
	}
	tokenString = strings.Replace(tokenString, "Bearer ", "", -1)

	return tokenString
}

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := getJwtToken(r)
		if tokenString == "" {
			ErrorJSONResponse(w, http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		ctx, err := ValidateToken(r.Context(), tokenString)
		if err != nil {
			ErrorJSONResponse(w, http.StatusUnauthorized, err)
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
