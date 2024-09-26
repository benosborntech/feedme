package businessserver

import (
	"github.com/benosborntech/feedme/pb"
)

type BusinessServer struct {
	pb.UnimplementedBusinessServer
}

func NewBusinessServer() *BusinessServer {
	return &BusinessServer{}
}
