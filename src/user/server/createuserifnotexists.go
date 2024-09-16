package userserver

import (
	"context"
	"fmt"

	"github.com/benosborntech/feedme/common/types"
	"github.com/benosborntech/feedme/pb"
	"github.com/benosborntech/feedme/user/dal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u *UserServer) CreateUserIfNotExists(ctx context.Context, req *pb.CreateUserIfNotExistsRequest) (*pb.CreateUserIfNotExistsResponse, error) {
	user, err := dal.CreateUser(ctx, u.db, &types.User{
		Id:    int(req.User.Id),
		Email: req.User.Email,
		Name:  req.User.Name,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to create user: %v", err))
	}

	res := &pb.CreateUserIfNotExistsResponse{
		User: &pb.UserData{
			Id:        int64(user.Id),
			Email:     user.Email,
			Name:      user.Name,
			UpdatedAt: timestamppb.New(user.UpdatedAt),
			CreatedAt: timestamppb.New(user.CreatedAt),
		},
	}

	return res, nil
}
