package handlers

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"

	"github.com/benosborntech/feedme/apigw/handlers/oauth"
	"github.com/benosborntech/feedme/apigw/types"
	commonTypes "github.com/benosborntech/feedme/common/types"
	"github.com/benosborntech/feedme/pb"
	"github.com/gofiber/fiber/v3"
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

func OAuthCallbackHandler(oauth oauth.OAuth, userClient pb.UserClient) func(c fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		log.Printf("oauthCallbackHandler.req=%v", string(c.Body()))

		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

		var body oauthCallbackHandlerBody
		if err := json.Unmarshal(c.Body(), &body); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("failed to unmarshal body, err=%v", err))
		}

		t, err := oauth.ExchangeToken(c.Context(), body.Code, body.State, body.Session)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("failed to get token, err=%v", err))
		}

		log.Printf("got token, token=%v", t)

		userInfo, err := oauth.GetUserInfo(c.Context(), t.AccessToken)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("failed to get user info, err=%v", err))
		}

		userId := getUserId(oauth.GetServiceType(), userInfo.Sub)

		// Create a user if they do not exist, or retrieve the existing one if they do
		var user commonTypes.User
		res, err := userClient.CreateUserIfNotExists(c.Context(), &pb.CreateUserIfNotExistsRequest{
			User: &pb.UserData{
				Id:    int64(userId),
				Email: userInfo.Email,
				Name:  userInfo.Name,
			},
		})
		if err != nil {
			res, err := userClient.GetUser(c.Context(), &pb.GetUserRequest{
				UserId: int64(userId),
			})
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("failed to get user, err=%v", err))
			}

			user = commonTypes.User{
				Id:        int(res.User.Id),
				Email:     res.User.Email,
				Name:      res.User.Name,
				UpdatedAt: res.User.UpdatedAt.AsTime(),
				CreatedAt: res.User.CreatedAt.AsTime(),
			}
		} else {
			user = commonTypes.User{
				Id:        int(res.User.Id),
				Email:     res.User.Email,
				Name:      res.User.Name,
				UpdatedAt: res.User.UpdatedAt.AsTime(),
				CreatedAt: res.User.CreatedAt.AsTime(),
			}
		}

		token := &types.Token{
			AccessToken:  t.AccessToken,
			RefreshToken: t.RefreshToken,
			TokenType:    oauth.GetServiceType(),
			User:         user,
			ExpiresAt:    t.Expiry,
		}
		payload, err := json.Marshal(token)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("failed to marshal token, err=%v", err))
		}

		return c.Status(fiber.StatusOK).SendString(string(payload))
	}
}
