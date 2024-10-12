package handlers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/benosborntech/feedme/common/types"
	"github.com/benosborntech/feedme/pb"
)

func GetUpdatesHandler(updatesClient pb.UpdatesClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		longX, err := strconv.ParseFloat(r.URL.Query().Get("longX"), 32)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get long x, err=%v", err), http.StatusBadRequest)
			return
		}

		latY, err := strconv.ParseFloat(r.URL.Query().Get("latY"), 32)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get lat y, err=%v", err), http.StatusBadRequest)
			return
		}

		radius, err := strconv.ParseFloat(r.URL.Query().Get("radius"), 32)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get radius, err=%v", err), http.StatusBadRequest)
			return
		}

		// Stream responses
		req := &pb.GetUpdatesRequest{
			LongX:  float32(longX),
			LatY:   float32(latY),
			Radius: float32(radius),
		}

		stream, err := updatesClient.GetUpdates(r.Context(), req)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get updates, err=%v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Transfer-Encoding", "chunked")

		writer := bufio.NewWriter(w)

		for {
			select {
			case <-r.Context().Done():
				log.Print("client disconnected")

				return
			default:
				update, err := stream.Recv()
				if err == io.EOF {
					log.Print("end of file")

					return
				}
				if err != nil {
					log.Printf("failed to receive data from stream, err=%v", err)

					continue
				}

				event := &types.Event{
					Item: types.Item{
						Id:         int(update.Item.Id),
						Location:   update.Item.Location,
						ItemType:   int(update.Item.ItemType),
						Quantity:   int(update.Item.Quantity),
						ExpiresAt:  update.Item.ExpiresAt.AsTime(),
						CreatedBy:  int(update.Item.CreatedBy),
						BusinessId: int(update.Item.BusinessId),
						UpdatedAt:  update.Item.UpdatedAt.AsTime(),
						CreatedAt:  update.Item.CreatedAt.AsTime(),
					},
				}

				payload, err := json.Marshal(event)
				if err != nil {
					log.Printf("failed to encode payload, err=%v", err)

					continue
				}

				fmt.Fprintln(writer, string(payload))
				if err := writer.Flush(); err != nil {
					log.Printf("failed to flush, err=%v", err)
				}
			}
		}
	}
}
