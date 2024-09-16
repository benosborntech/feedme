package handlers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/benosborntech/feedme/apigw/handlers/oauth"
	"github.com/gofiber/fiber/v3"
)

type handlerResponse struct {
	RedirectEndpoint string `json:"redirectEndpoint"`
	Session          string `json:"session"`
}

func GetOAuthEndpointHandler(oauth oauth.OAuth) func(c fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		log.Printf("getOAuthEndpointHandleretOAuthHandler.req=%v", string(c.Body()))

		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

		session, endpoint, err := oauth.GetEndpoint(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("failed to get redirect endpoint, err=%v", err))
		}

		response := &handlerResponse{
			RedirectEndpoint: endpoint,
			Session:          session,
		}
		payload, err := json.Marshal(response)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("failed to marshal response, err=%v", err))
		}

		return c.Status(fiber.StatusOK).SendString(string(payload))
	}
}
