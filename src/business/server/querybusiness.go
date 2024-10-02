package businessserver

import (
	"context"
	"fmt"

	"github.com/benosborntech/feedme/business/dal"
	"github.com/benosborntech/feedme/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (b *BusinessServer) QueryBusiness(ctx context.Context, req *pb.QueryBusinessRequest) (*pb.QueryBusinessResponse, error) {
	businesses, err := dal.QueryBusiness(ctx, b.db, int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to get business: %v", err))
	}

	out := []*pb.BusinessData{}

	for _, business := range businesses {
		out = append(out, &pb.BusinessData{
			Id:          int64(business.Id),
			Name:        business.Name,
			Description: business.Description,
			Latitude:    business.Latitude,
			Longitude:   business.Longitude,
			CreatedBy:   int64(business.CreatedBy),
			UpdatedAt:   timestamppb.New(business.UpdatedAt),
			CreatedAt:   timestamppb.New(business.CreatedAt),
		})
	}

	return &pb.QueryBusinessResponse{
		Business: out,
	}, nil
}
