package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/benosborntech/feedme/pb"
	"github.com/benosborntech/feedme/user/config"
	userserver "github.com/benosborntech/feedme/user/server"
	_ "github.com/go-sql-driver/mysql"
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

	userServer := userserver.NewUserServer(db)

	tcpListener, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		log.Fatalf("failed to start tcp listener, err=%v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterUserServer(s, userServer)
	if err := s.Serve(tcpListener); err != nil {
		log.Fatalf("failed to serve, err=%v", err)
	}
}
