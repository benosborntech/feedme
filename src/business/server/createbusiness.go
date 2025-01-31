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
	business, err := dal.CreateBusiness(ctx, b.db, &types.Business{
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
			Id:          int64(business.Id),
			Name:        business.Name,
			Description: business.Description,
			Latitude:    business.Latitude,
			Longitude:   business.Longitude,
			CreatedBy:   int64(business.CreatedBy),
			UpdatedAt:   timestamppb.New(business.UpdatedAt),
			CreatedAt:   timestamppb.New(business.CreatedAt),
		},
	}, nil
}
