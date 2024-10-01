package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/benosborntech/feedme/item/config"
	"github.com/benosborntech/feedme/item/poller"
	itemserver "github.com/benosborntech/feedme/item/server"
	"github.com/benosborntech/feedme/pb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.NewConfig()

	db, err := sql.Open("mysql", cfg.MySQLDSN)
	if err != nil {
		log.Fatalf("failed to open mysql connection, err=%v", err)
	}
	defer db.Close()

	client := redis.NewClient(cfg.RedisCfg)
	defer client.Close()

	ctx := context.Background()

	poll := poller.NewPoller(db, client)
	go func(ctx context.Context) {
		if err := poll.Poll(ctx); err != nil {
			log.Fatalf("poller failed, err=%v", err)
		}
	}(ctx)

	itemServer := itemserver.NewItemServer(db)

	tcpListener, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		log.Fatalf("failed to start tcp listener, err=%v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterItemServer(s, itemServer)
	if err := s.Serve(tcpListener); err != nil {
		log.Fatalf("failed to serve, err=%v", err)
	}
}
