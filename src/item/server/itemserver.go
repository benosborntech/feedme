package itemserver

import (
	"database/sql"

	"github.com/benosborntech/feedme/pb"
)

type ItemServer struct {
	pb.UnimplementedItemServer
	db *sql.DB
}

func NewItemServer(db *sql.DB) *ItemServer {
	return &ItemServer{
		db: db,
	}
}
