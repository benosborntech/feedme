package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/benosborntech/feedme/apigw/config"
	"github.com/benosborntech/feedme/apigw/oauth"
	"github.com/benosborntech/feedme/apigw/types"
	"github.com/redis/go-redis/v9"
)

type refreshTokenHandlerRequestBody struct {
	RefreshToken string            `json:"refreshToken"`
	TokenType    types.ServiceType `json:"tokenType"`
}

type refreshTokenHandlerResponseBody struct {
	*types.Token
}

func RefreshTokenHandler(cfg *config.Config, client *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body refreshTokenHandlerRequestBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, fmt.Sprintf("failed to unmarshal body, err=%v", err), http.StatusInternalServerError)
			return
		}

		var oauthHandler oauth.OAuth
		switch body.TokenType {
		case types.GoogleType:
			oauthHandler = oauth.NewOAuthGoogle(client, cfg.GoogleOAuthConfig, cfg.BaseURL)
		default:
			http.Error(w, fmt.Sprintf("invalid token type, token type=%v", body.TokenType), http.StatusBadRequest)
			return
		}

		token, err := oauthHandler.RefreshToken(r.Context(), body.RefreshToken)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to refresh token, err=%v", err), http.StatusInternalServerError)
			return
		}

		// **** Here I will generate a new refresh token for the user, fill it out with the required details, and do the rest
		// **** Might come up with a function for creating the JWT if possible...

		response := &refreshTokenHandlerResponseBody{
			Token: &types.Token{},
		}
	}
}
