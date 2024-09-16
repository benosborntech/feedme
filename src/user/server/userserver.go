package userserver

import (
	"database/sql"

	"github.com/benosborntech/feedme/pb"
)

type UserServer struct {
	pb.UnimplementedUserServer
	db *sql.DB
}

func NewUserServer(db *sql.DB) *UserServer {
	return &UserServer{
		db: db,
	}
}
