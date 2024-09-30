package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/benosborntech/feedme/apigw/config"
	"github.com/benosborntech/feedme/apigw/oauth"
	"github.com/benosborntech/feedme/apigw/utils"
	"github.com/benosborntech/feedme/pb"
)

type oauthCallbackHandlerRequestBody struct {
	Code    string `json:"code"`
	State   string `json:"state"`
	Session string `json:"session"`
}

func OAuthCallbackHandler(cfg *config.Config, oauth oauth.OAuth, userClient pb.UserClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body oauthCallbackHandlerRequestBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, fmt.Sprintf("failed to unmarshal body, err=%v", err), http.StatusInternalServerError)
			return
		}

		t, err := oauth.ExchangeToken(r.Context(), body.Code, body.State, body.Session)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get token, err=%v", err), http.StatusInternalServerError)
			return
		}

		userInfo, err := oauth.GetUserInfo(r.Context(), t.AccessToken)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get user info, err=%v", err), http.StatusInternalServerError)
			return
		}

		refreshToken := t.RefreshToken
		tokenType := oauth.GetServiceType()
		userId := utils.GetUserId(oauth.GetServiceType(), userInfo.Sub)
		expiresAt := t.Expiry

		_, err = userClient.CreateUserIfNotExists(r.Context(), &pb.CreateUserIfNotExistsRequest{
			User: &pb.UserData{
				Id:    int64(userId),
				Email: userInfo.Email,
				Name:  userInfo.Name,
			},
		})
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get user, err=%v", err), http.StatusInternalServerError)
			return
		}

		TokenHandler(cfg, w, r, &TokenHandlerRequestBody{
			RefreshToken: refreshToken,
			TokenType:    tokenType,
			UserId:       userId,
			ExpiresAt:    expiresAt,
		})
	}
}
