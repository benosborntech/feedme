package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/benosborntech/feedme/common/types"
	"github.com/benosborntech/feedme/pb"
)

type getBusinessesHandlerRequestBody struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

type getBusinessesHandlerResponseBody struct {
	Business []*types.Business `json:"business"`
}

func GetBusinessesHandler(businessClient pb.BusinessClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body getBusinessesHandlerRequestBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, fmt.Sprintf("failed to unmarshal body, err=%v", err), http.StatusInternalServerError)
			return
		}

		data, err := businessClient.QueryBusiness(r.Context(), &pb.QueryBusinessRequest{
			Page:     int32(body.Page),
			PageSize: int32(body.PageSize),
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

type postBusinessesHandlerRequestBody struct {
	Business *types.Business `json:"business"`
}

type postBusinessesHandlerResponseBody struct {
	Business *types.Business `json:"business"`
}

func GetBusinessesHandler(businessClient pb.BusinessClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body postBusinessesHandlerRequestBody
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
				CreatedBy:   int64(body.Business.CreatedBy),
			},
		})
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to query business, err=%v", err), http.StatusInternalServerError)
			return
		}

		businesses := []*types.Business{}

		// **** Ok, I need to come up with a better way of protecting my endpoints than this approach - this is currently not going to work. I need to add my own type of routing layer here...

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

func BusinessesHandler(businessClient pb.BusinessClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			GetBusinessesHandler(businessClient)(w, r)
		} else if r.Method == http.MethodPost {

		} else {
			http.Error(w, "method does not exist", http.StatusNotFound)
			return
		}
	}
}
