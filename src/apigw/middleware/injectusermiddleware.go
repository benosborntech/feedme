package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/benosborntech/feedme/apigw/config"
	"github.com/benosborntech/feedme/apigw/types"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

func InjectUserMiddleware(cfg *config.Config, client *redis.Client, next func(userId int) http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			http.Error(w, "no authorization header", http.StatusUnauthorized)
			return
		}

		split := strings.Split(authorization, " ")
		if len(split) != 2 || split[0] != "Bearer" {
			http.Error(w, "invalid authorization header format", http.StatusBadRequest)
			return
		}

		jwtToken := split[1]
		parsedJwtToken, err := jwt.ParseWithClaims(jwtToken, &types.Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(cfg.ServerSecret), nil
		})
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to parse token, err=%v", err), http.StatusBadRequest)
			return
		}
		claims, ok := parsedJwtToken.Claims.(*types.Claims)
		if !ok || !parsedJwtToken.Valid {
			http.Error(w, fmt.Sprintf("token is not valid, err=%v", err), http.StatusUnauthorized)
			return
		}

		userId, err := strconv.Atoi(claims.RegisteredClaims.Subject)
		if err != nil {
			http.Error(w, fmt.Sprintf("could not get user id, err=%v", err), http.StatusInternalServerError)
			return
		}

		// Execute the handler
		next(userId)(w, r)
	}
}
