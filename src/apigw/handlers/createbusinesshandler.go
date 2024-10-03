package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/benosborntech/feedme/common/types"
	"github.com/benosborntech/feedme/pb"
)

type createBusinessHandlerRequestBody struct {
	Business *types.Business `json:"business"`
}

type createBusinessHandlerResponseBody struct {
	Business *types.Business `json:"business"`
}

func CreateBusinessesHandler(businessClient pb.BusinessClient) func(userId int) http.HandlerFunc {
	return func(userId int) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var body createBusinessHandlerRequestBody
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				http.Error(w, fmt.Sprintf("failed to unmarshal body, err=%v", err), http.StatusInternalServerError)
				return
			}

			data, err := businessClient.CreateBusiness(r.Context(), &pb.CreateBusinessRequest{
				Business: &pb.BusinessData{
					Name:        body.Business.Name,
					Description: body.Business.Description,
					Latitude:    body.Business.Latitude,
					Longitude:   body.Business.Longitude,
					CreatedBy:   int64(userId),
				},
			})
			if err != nil {
				http.Error(w, fmt.Sprintf("failed to query business, err=%v", err), http.StatusInternalServerError)
				return
			}

			response := &createBusinessHandlerResponseBody{
				Business: &types.Business{
					Id:          int(data.Business.Id),
					Name:        data.Business.Name,
					Description: data.Business.Description,
					Latitude:    data.Business.Latitude,
					Longitude:   data.Business.Longitude,
					CreatedBy:   int(data.Business.CreatedBy),
					UpdatedAt:   data.Business.UpdatedAt.AsTime(),
					CreatedAt:   data.Business.CreatedAt.AsTime(),
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
