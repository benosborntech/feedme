package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/benosborntech/feedme/common/types"
	"github.com/benosborntech/feedme/pb"
)

type createItemHandlerRequestBody struct {
	Item *types.Item `json:"item"`
}

type createItemHandlerResponseBody struct {
	Item *types.Item `json:"item"`
}

func CreateItemHandler(itemClient pb.ItemClient) func(userId int) http.HandlerFunc {
	return func(userId int) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var body createItemHandlerRequestBody
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				http.Error(w, fmt.Sprintf("failed to unmarshal body, err=%v", err), http.StatusInternalServerError)
				return
			}

			// **** We probably need to provide some data from the business here to fill this out properly...

			data, err := itemClient.CreateItem(r.Context(), &pb.CreateItemRequest{
				Item: &pb.ItemData{
					// **** Add the data in here...
					Location: body.Item.Location,
				},
			})
			if err != nil {
				http.Error(w, fmt.Sprintf("failed to create item, err=%v", err), http.StatusInternalServerError)
				return
			}

			response := &createItemHandlerResponseBody{
				Item: &types.Item{
					// **** Add the response in here
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
	}
}
