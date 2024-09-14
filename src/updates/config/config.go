package config

import (
	"log"
	"os"

	"github.com/benosborntech/feedme/updates/consts"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Port     string
	RedisCfg *redis.Options
}

func NewConfig() *Config {
	port := consts.PORT
	if envPort, ok := os.LookupEnv("PORT"); ok {
		port = envPort

		log.Printf("using custom port, port=%v", port)
	} else {
		log.Printf("using default port, port=%v", port)
	}

	redisAddr, ok := os.LookupEnv("REDIS_ADDR")
	if !ok {
		log.Print("no redis address provided")
	} else {
		log.Printf("using redis addr, addr=%v", redisAddr)
	}

	return &Config{
		Port: port,
		RedisCfg: &redis.Options{
			Addr: redisAddr,
		},
	}
}
