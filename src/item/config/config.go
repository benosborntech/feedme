package config

import (
	"log"
	"os"

	"github.com/benosborntech/feedme/apigw/consts"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Port     string
	MySQLDSN string
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

	mySQLDSN, ok := os.LookupEnv("MYSQL_DSN")
	if !ok {
		log.Printf("no mysql dsn provided")
	} else {
		log.Printf("using mysql dsn, dsn=%v", mySQLDSN)
	}

	redisAddr, ok := os.LookupEnv("REDIS_ADDR")
	if !ok {
		log.Print("no redis addr provided")
	} else {
		log.Printf("using redis addr, addr=%v", redisAddr)
	}

	return &Config{
		Port:     port,
		MySQLDSN: mySQLDSN,
		RedisCfg: &redis.Options{
			Addr: redisAddr,
		},
	}
}
