package itemserver

import (
	"context"
	"fmt"

	"github.com/benosborntech/feedme/item/dal"
	"github.com/benosborntech/feedme/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (i *ItemServer) GetItem(ctx context.Context, req *pb.GetItemRequest) (*pb.GetItemResponse, error) {
	item, err := dal.QueryItemById(ctx, i.db, int(req.ItemId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to get item: %v", err))
	}

	return &pb.GetItemResponse{
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
