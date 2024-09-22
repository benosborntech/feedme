package config

import (
	"log"
	"os"

	"github.com/benosborntech/feedme/apigw/types"
	"github.com/benosborntech/feedme/updates/consts"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Port              string
	UpdatesAddr       string
	UserAddr          string
	BaseURL           string
	ServerSecret      string
	RedisCfg          *redis.Options
	GoogleOAuthConfig *types.OAuthConfig
}

func NewConfig() *Config {
	port := consts.PORT
	if envPort, ok := os.LookupEnv("PORT"); ok {
		port = envPort

		log.Printf("using custom port, port=%v", port)
	} else {
		log.Printf("using default port, port=%v", port)
	}

	updatesAddr, ok := os.LookupEnv("UPDATES_ADDR")
	if !ok {
		log.Print("no updates addr provided")
	} else {
		log.Printf("using updates addr, addr=%v", updatesAddr)
	}

	userAddr, ok := os.LookupEnv("USER_ADDR")
	if !ok {
		log.Print("no user addr provided")
	} else {
		log.Printf("using user addr, addr=%v", userAddr)
	}

	redisAddr, ok := os.LookupEnv("REDIS_ADDR")
	if !ok {
		log.Print("no redis addr provided")
	} else {
		log.Printf("using redis addr, addr=%v", redisAddr)
	}

	baseURL, ok := os.LookupEnv("BASE_URL")
	if !ok {
		log.Print("no base url provided")
	} else {
		log.Printf("using base url, base url=%v", baseURL)
	}

	googleClientId, ok := os.LookupEnv("GOOGLE_CLIENT_ID")
	if !ok {
		log.Print("no google client id provided")
	} else {
		log.Printf("using google client id, id=%v***...***", googleClientId[:3])
	}

	googleClientSecret, ok := os.LookupEnv("GOOGLE_CLIENT_SECRET")
	if !ok {
		log.Print("no google client secret provided")
	} else {
		log.Printf("using google client secret, id=%v***...***", googleClientSecret[:3])
	}

	serverSecret, ok := os.LookupEnv("SERVER_SECRET")
	if !ok {
		log.Print("no server secret provided")
	} else {
		log.Printf("using server secret, id=%v***...***", serverSecret[:3])
	}

	return &Config{
		Port:         port,
		UpdatesAddr:  updatesAddr,
		UserAddr:     userAddr,
		BaseURL:      baseURL,
		ServerSecret: serverSecret,
		RedisCfg: &redis.Options{
			Addr: redisAddr,
		},
		GoogleOAuthConfig: &types.OAuthConfig{
			ClientId:     googleClientId,
			ClientSecret: googleClientSecret,
		},
	}
}
