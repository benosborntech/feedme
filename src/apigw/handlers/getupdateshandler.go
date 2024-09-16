package handlers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/benosborntech/feedme/common/types"
	"github.com/benosborntech/feedme/pb"
	"github.com/gofiber/fiber/v3"
	"github.com/valyala/fasthttp"
)

type updatesBody struct {
	LongX  float32 `json:"longX"`
	LatY   float32 `json:"latY"`
	Radius float32 `json:"radius"`
}

func GetUpdatesHandler(updatesClient pb.UpdatesClient) func(c fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		log.Printf("getUpdatesHandler.req=%v", string(c.Body()))

		var body updatesBody
		if err := json.Unmarshal(c.Body(), &body); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("failed to parse request, err=%v", err))
		}

		c.Set("Content-Type", "text/event-stream")
		c.Set("Cache-Control", "no-cache")
		c.Set("Connection", "keep-alive")
		c.Set("Transfer-Encoding", "chunked")

		// Stream responses
		req := &pb.GetUpdatesRequest{
			LongX:  body.LongX,
			LatY:   body.LatY,
			Radius: body.Radius,
		}

		stream, err := updatesClient.GetUpdates(c.Context(), req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("failed to get updates, err=%v", err))
		}

		c.Status(fiber.StatusOK).Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
			for {
				select {
				case <-c.Context().Done():
					return
				default:
					update, err := stream.Recv()
					if err == io.EOF {
						return
					}
					if err != nil {
						log.Printf("failed to receive data from stream, err=%v", err)

						continue
					}

					event := &types.Event{
						Item: types.Item{
							Id:        int(update.Item.Id),
							Location:  update.Item.Location,
							ItemType:  int(update.Item.ItemType),
							Quantity:  int(update.Item.Quantity),
							UpdatedAt: update.Item.UpdatedAt.AsTime(),
							CreatedAt: update.Item.CreatedAt.AsTime(),
						},
					}
					payload, err := json.Marshal(event)
					if err != nil {
						log.Printf("failed to encode payload, err=%v", err)

						continue
					}

					fmt.Fprint(w, string(payload))
					if err := w.Flush(); err != nil {
						log.Printf("failed to flush, err=%v", err)
					}
				}
			}
		}))

		return c.Status(fiber.StatusOK).SendString("closed connection")
	}
}
