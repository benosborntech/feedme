package main

import (
	"fmt"
	"log"
	"net"

	"github.com/benosborntech/feedme/pb"
	"github.com/benosborntech/feedme/updates/config"
	"github.com/benosborntech/feedme/updates/listener"
	updatesserver "github.com/benosborntech/feedme/updates/server"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.NewConfig()

	client := redis.NewClient(cfg.RedisCfg)
	defer client.Close()

	listener := listener.NewListener(client)

	updatesServer := updatesserver.NewUpdatesServer(listener)

	tcpListener, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		log.Fatalf("failed to start tcp listener, err=%v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterUpdatesServer(s, updatesServer)
	if err := s.Serve(tcpListener); err != nil {
		log.Fatalf("failed to serve, err=%v", err)
	}
}
