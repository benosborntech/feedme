package handlers

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/benosborntech/feedme/apigw/oauth"
	"github.com/benosborntech/feedme/apigw/types"
	"github.com/benosborntech/feedme/pb"
)

type oauthCallbackHandlerBody struct {
	Code    string `json:"code"`
	State   string `json:"state"`
	Session string `json:"session"`
}

func getUserId(serviceType types.ServiceType, sub string) int {
	uniqueId := fmt.Sprintf("%s:%s", serviceType, sub)

	hash := sha256.Sum256([]byte(uniqueId))
	userId := binary.BigEndian.Uint64(hash[:8])

	return int(userId)
}

func OAuthCallbackHandler(oauth oauth.OAuth, userClient pb.UserClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body oauthCallbackHandlerBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, fmt.Sprintf("failed to unmarshal body, err=%v", err), http.StatusInternalServerError)
			return
		}

		t, err := oauth.ExchangeToken(r.Context(), body.Code, body.State, body.Session)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get token, err=%v", err), http.StatusInternalServerError)
			return
		}

		log.Printf("got token, token=%v", t)

		userInfo, err := oauth.GetUserInfo(r.Context(), t.AccessToken)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get user info, err=%v", err), http.StatusInternalServerError)
			return
		}

		userId := getUserId(oauth.GetServiceType(), userInfo.Sub)

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

		token := &types.Token{
			AccessToken:  t.AccessToken,
			RefreshToken: t.RefreshToken,
			TokenType:    oauth.GetServiceType(),
			UserId:       userId,
			ExpiresAt:    t.Expiry,
		}
		payload, err := json.Marshal(token)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to marshal token, err=%v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(payload)
	}
}
