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

func (b *BusinessServer) GetBusiness(ctx context.Context, req *pb.GetBusinessRequest) (*pb.GetBusinessResponse, error) {
	item, err := dal.QueryBusinessById(ctx, b.db, int(req.BusinessId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to get business: %v", err))
	}

	return &pb.GetBusinessResponse{
		Business: &pb.BusinessData{
			Id:          int64(item.Id),
			Name:        item.Name,
			Description: item.Description,
			Latitude:    item.Latitude,
			Longitude:   item.Longitude,
			CreatedBy:   int64(item.CreatedBy),
			UpdatedAt:   timestamppb.New(item.UpdatedAt),
			CreatedAt:   timestamppb.New(item.CreatedAt),
		},
	}, nil
}
