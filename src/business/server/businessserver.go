package businessserver

import (
	"database/sql"

	"github.com/benosborntech/feedme/pb"
)

type BusinessServer struct {
	pb.UnimplementedBusinessServer
	db *sql.DB
}

func NewBusinessServer(db *sql.DB) *BusinessServer {
	return &BusinessServer{
		db: db,
	}
}
