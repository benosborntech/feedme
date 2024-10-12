package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/benosborntech/feedme/common/types"
	"github.com/benosborntech/feedme/pb"
)

type getBusinessesHandlerResponseBody struct {
	Business []*types.Business `json:"business"`
}

func GetBusinessesHandler(businessClient pb.BusinessClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get page, err=%v", err), http.StatusBadRequest)
			return
		}

		pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get page size, err=%v", err), http.StatusBadRequest)
			return
		}

		data, err := businessClient.QueryBusiness(r.Context(), &pb.QueryBusinessRequest{
			Page:     int32(page),
			PageSize: int32(pageSize),
		})
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to query business, err=%v", err), http.StatusInternalServerError)
			return
		}

		businesses := []*types.Business{}

		for _, business := range data.Business {
			businesses = append(businesses, &types.Business{
				Id:          int(business.Id),
				Name:        business.Name,
				Description: business.Description,
				Latitude:    business.Latitude,
				Longitude:   business.Longitude,
				CreatedBy:   int(business.CreatedBy),
				UpdatedAt:   business.UpdatedAt.AsTime(),
				CreatedAt:   business.CreatedAt.AsTime(),
			})
		}

		response := &getBusinessesHandlerResponseBody{
			Business: businesses,
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
