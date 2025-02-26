package auth

import (
	"context"
	"net/http"
	"os"

	"github.com/go-chi/jwtauth/v5"
)

var TokenAuth *jwtauth.JWTAuth

func init() {
	secret := os.Getenv("JWT_SECRET")
	TokenAuth = jwtauth.New("HS256", []byte(secret), nil)
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		userID, ok := claims["user_id"].(string)
		if !ok {
			http.Error(w, "Token error", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GenerateToken(userID string) (string, error) {
	_, tokenString, err := TokenAuth.Encode(map[string]interface{}{"user_id": userID})
	return tokenString, err
}
