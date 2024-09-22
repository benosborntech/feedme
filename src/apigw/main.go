package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/benosborntech/feedme/apigw/config"
	"github.com/benosborntech/feedme/apigw/handlers"
	"github.com/benosborntech/feedme/apigw/handlers/oauth"
	"github.com/benosborntech/feedme/pb"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg := config.NewConfig()

	client := redis.NewClient(cfg.RedisCfg)
	defer client.Close()

	updatesConn, err := grpc.NewClient(cfg.UpdatesAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to start updates conn, err=%v", err)
	}
	defer updatesConn.Close()
	updatesClient := pb.NewUpdatesClient(updatesConn)

	userConn, err := grpc.NewClient(cfg.UpdatesAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to start updates conn, err=%v", err)
	}
	defer userConn.Close()
	userClient := pb.NewUserClient(userConn)

	// Register OAuth endpoints
	oauthHandlers := []oauth.OAuth{
		oauth.NewOAuthGoogle(client, cfg.GoogleOAuthConfig, cfg.BaseURL),
	}
	for _, handler := range oauthHandlers {
		http.HandleFunc(handler.GetEndpointPath(), handlers.GetOAuthEndpointHandler(handler))
		http.HandleFunc(handler.GetCallbackPath(), handlers.OAuthCallbackHandler(handler, userClient))
	}

	http.HandleFunc("/api/updates", handlers.GetUpdatesHandler(updatesClient))

	log.Printf("started server, addr=http://localhost:%s", cfg.Port)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), nil); err != nil {
		log.Fatalf("fatal server error, err=%v", err)
	}
}
