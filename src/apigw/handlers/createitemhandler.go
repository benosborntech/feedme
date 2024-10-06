package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/benosborntech/feedme/common/consts"
	"github.com/benosborntech/feedme/common/types"
	"github.com/benosborntech/feedme/pb"
	"github.com/pierrre/geohash"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type createItemHandlerRequestBody struct {
	Item     *types.Item     `json:"item"`
	Business *types.Business `json:"business"`
}

type createItemHandlerResponseBody struct {
	Item *types.Item `json:"item"`
}

func CreateItemHandler(itemClient pb.ItemClient, businessClient pb.BusinessClient) func(userId int) http.HandlerFunc {
	return func(userId int) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var body createItemHandlerRequestBody
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				http.Error(w, fmt.Sprintf("failed to unmarshal body, err=%v", err), http.StatusInternalServerError)
				return
			}

			// Lookup the business
			business, err := businessClient.GetBusiness(r.Context(), &pb.GetBusinessRequest{
				BusinessId: int64(body.Business.Id),
			})
			if err != nil {
				http.Error(w, fmt.Sprintf("failed to get business, err=%v", err), http.StatusInternalServerError)
				return
			}

			// Calculate geo hash using maximum precision to ensure all listeners listening to the chunk are notified
			latY, err := strconv.ParseFloat(business.Business.Latitude, 64)
			if err != nil {
				http.Error(w, fmt.Sprintf("failed to parse latitude, err=%v", err), http.StatusInternalServerError)
				return
			}
			longX, err := strconv.ParseFloat(business.Business.Longitude, 64)
			if err != nil {
				http.Error(w, fmt.Sprintf("failed to parse longitude, err=%v", err), http.StatusInternalServerError)
				return
			}

			location := geohash.Encode(float64(latY), float64(longX), consts.MAX_PRECISION)

			item, err := itemClient.CreateItem(r.Context(), &pb.CreateItemRequest{
				Item: &pb.ItemData{
					Location:   location,
					ItemType:   int32(body.Item.ItemType),
					Quantity:   int32(body.Item.Quantity),
					ExpiresAt:  timestamppb.New(body.Item.ExpiresAt),
					CreatedBy:  int64(userId),
					BusinessId: business.Business.Id,
				},
			})
			if err != nil {
				http.Error(w, fmt.Sprintf("failed to create item, err=%v", err), http.StatusInternalServerError)
				return
			}

			response := &createItemHandlerResponseBody{
				Item: &types.Item{
					Id:         int(item.Item.Id),
					Location:   item.Item.Location,
					ItemType:   int(item.Item.ItemType),
					Quantity:   int(item.Item.Quantity),
					ExpiresAt:  item.Item.ExpiresAt.AsTime(),
					CreatedBy:  int(item.Item.CreatedBy),
					BusinessId: int(item.Item.BusinessId),
					UpdatedAt:  item.Item.UpdatedAt.AsTime(),
					CreatedAt:  item.Item.CreatedAt.AsTime(),
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
