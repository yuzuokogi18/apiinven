package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

var UserContextKey = contextKey("user")

type Claims struct {
	UserID int
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Token requerido", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		// 🔥 Forma segura (sin panic)
		claimsMap, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		userIDFloat, ok := claimsMap["user_id"].(float64)
		if !ok {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		userID := int(userIDFloat)

		claims := &Claims{UserID: userID}
		ctx := context.WithValue(r.Context(), UserContextKey, claims)

		next(w, r.WithContext(ctx))
	}
}
