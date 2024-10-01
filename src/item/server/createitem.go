package itemserver

import (
	"context"
	"fmt"

	"github.com/benosborntech/feedme/common/types"
	"github.com/benosborntech/feedme/item/dal"
	"github.com/benosborntech/feedme/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (i *ItemServer) CreateItem(ctx context.Context, req *pb.CreateItemRequest) (*pb.CreateItemResponse, error) {
	item, err := dal.CreateItem(ctx, i.db, &types.Item{
		Location:   req.Item.Location,
		ItemType:   int(req.Item.ItemType),
		Quantity:   int(req.Item.Quantity),
		ExpiresAt:  req.Item.ExpiresAt.AsTime(),
		CreatedBy:  int(req.Item.CreatedBy),
		BusinessId: int(req.Item.BusinessId),
		UpdatedAt:  req.Item.UpdatedAt.AsTime(),
		CreatedAt:  req.Item.CreatedAt.AsTime(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to create item: %v", err))
	}

	return &pb.CreateItemResponse{
		Item: &pb.ItemData{
			Id:         int64(item.Id),
			Location:   item.Location,
			ItemType:   int32(item.ItemType),
			Quantity:   int32(item.Quantity),
			ExpiresAt:  timestamppb.New(item.ExpiresAt),
			CreatedBy:  int64(item.CreatedBy),
			BusinessId: int64(item.BusinessId),
			UpdatedAt:  timestamppb.New(item.UpdatedAt),
			CreatedAt:  timestamppb.New(item.CreatedAt),
		},
	}, nil
}
