package updatesserver

import (
	"context"
	"fmt"
	"log"
	"math"

	"github.com/benosborntech/feedme/common/consts"
	"github.com/benosborntech/feedme/common/types"
	"github.com/benosborntech/feedme/common/utils"
	"github.com/benosborntech/feedme/pb"
	"github.com/pierrre/geohash"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *UpdatesServer) GetUpdates(req *pb.GetUpdatesRequest, stream grpc.ServerStreamingServer[pb.GetUpdatesResponse]) error {
	log.Printf("itemsHandler.req=%v", req)

	radius := math.Min(consts.MAX_RADIUS, math.Max(consts.MIN_RADIUS, float64(req.Radius)))
	log.Printf("radius=%v", radius)

	// Calculate geo hash
	// Because we can use variable length precision here, we will need to publish to the topic of all hash codes
	// We also need to cap the maximum precision, so we make sure for all precisions here, we will always be publishing to it
	precision := utils.HashPrecision(radius)
	hash := geohash.Encode(float64(req.LatY), float64(req.LongX), precision)
	neighbors, err := geohash.GetNeighbors(hash)
	if err != nil {
		return fmt.Errorf("failed to get neighbors: %w", err)
	}

	hashes := []string{
		hash,
		neighbors.North,
		neighbors.South,
		neighbors.East,
		neighbors.West,
		neighbors.NorthEast,
		neighbors.NorthWest,
		neighbors.SouthEast,
		neighbors.SouthWest,
	}

	aggregatedCh := make(chan *types.Event)

	for _, hash := range hashes {
		id, ch, err := s.listener.Subscribe(hash)
		if err != nil {
			log.Printf("failed to subscribe to channel, channel=%v, err=%v", hash, err)

			continue
		}

		go func(ctx context.Context, aggregatedCh chan *types.Event, ch chan *types.Event) {
			defer func() {
				err := s.listener.Unsubscribe(hash, id)
				if err != nil {
					log.Printf("failed to unsubscribe from channel, channel=%v, err=%v", hash, err)
				}
			}()

			for {
				select {
				case msg := <-ch:
					aggregatedCh <- msg
				case <-ctx.Done():
					log.Printf("exiting aggregation channel, channel=%v, err=%v", hash, err)

					return
				}
			}
		}(stream.Context(), aggregatedCh, ch)
	}

	for {
		select {
		case msg := <-aggregatedCh:
			res := &pb.GetUpdatesResponse{
				Id:        int64(msg.Item.Id),
				Location:  msg.Item.Location,
				ItemType:  int32(msg.Item.ItemType),
				Quantity:  int32(msg.Item.Quantity),
				CreatedAt: timestamppb.New(msg.Item.CreatedAt),
			}
			if err := stream.Send(res); err != nil {
				log.Printf("failed to send stream response: res=%v, err=%v", res, err)

				continue
			}

			log.Printf("sent stream response: res=%v", res)
		case <-stream.Context().Done():
			log.Printf("client disconnected")

			return nil
		}
	}
}
