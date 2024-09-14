package updatesserver

import (
	"github.com/benosborntech/feedme/pb"
	"github.com/benosborntech/feedme/updates/listener"
)

type UpdatesServer struct {
	pb.UnimplementedUpdatesServer
	listener *listener.Listener
}

func NewUpdatesServer(listener *listener.Listener) *UpdatesServer {
	return &UpdatesServer{
		listener: listener,
	}
}
