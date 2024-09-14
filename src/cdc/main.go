package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/benosborntech/feedme/cdc/config"
	"github.com/benosborntech/feedme/cdc/poller"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
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

	<-ctx.Done()
}
