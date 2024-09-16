package userserver

import (
	"context"
	"fmt"

	"github.com/benosborntech/feedme/pb"
	"github.com/benosborntech/feedme/user/dal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := dal.GetUserByUserId(ctx, u.db, int(req.UserId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to get user: %v", err))
	}
	if user == nil {
		return nil, status.Errorf(codes.InvalidArgument, "user does not exist")
	}

	res := &pb.GetUserResponse{
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
