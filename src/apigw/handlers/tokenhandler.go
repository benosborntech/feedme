package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/benosborntech/feedme/apigw/config"
	"github.com/benosborntech/feedme/apigw/consts"
	"github.com/benosborntech/feedme/apigw/types"
	"github.com/golang-jwt/jwt/v5"
)

type TokenHandlerRequestBody struct {
	RefreshToken string
	TokenType    types.ServiceType
	UserId       int
	ExpiresAt    time.Time
}

type tokenHandlerResponseBody struct {
	*types.Token
}

func TokenHandler(cfg *config.Config, w http.ResponseWriter, r *http.Request, params *TokenHandlerRequestBody) {
	claims := types.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    consts.JWT_ISSUER,
			Subject:   fmt.Sprint(params.UserId),
			ExpiresAt: jwt.NewNumericDate(params.ExpiresAt),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(cfg.ServerSecret))
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to sign token, err=%v", err), http.StatusInternalServerError)
		return
	}

	response := &tokenHandlerResponseBody{
		Token: &types.Token{
			AccessToken:  token,
			RefreshToken: params.RefreshToken,
			TokenType:    params.TokenType,
			UserId:       params.UserId,
			ExpiresAt:    params.ExpiresAt,
		},
	}
	payload, err := json.Marshal(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal response, err=%v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}
