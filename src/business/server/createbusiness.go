package businessserver

import (
	"context"
	"fmt"

	"github.com/benosborntech/feedme/business/dal"
	"github.com/benosborntech/feedme/common/types"
	"github.com/benosborntech/feedme/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (b *BusinessServer) CreateBusiness(ctx context.Context, req *pb.CreateBusinessRequest) (*pb.CreateBusinessResponse, error) {
	item, err := dal.CreateBusiness(ctx, b.db, &types.Business{
		Name:        req.Business.Name,
		Description: req.Business.Description,
		Latitude:    req.Business.Latitude,
		Longitude:   req.Business.Longitude,
		CreatedBy:   int(req.Business.CreatedBy),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to create business: %v", err))
	}

	return &pb.CreateBusinessResponse{
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
