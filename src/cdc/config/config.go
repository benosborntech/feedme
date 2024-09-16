package config

import (
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	MySQLDSN string
	RedisCfg *redis.Options
}

func NewConfig() *Config {
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
		MySQLDSN: mySQLDSN,
		RedisCfg: &redis.Options{
			Addr: redisAddr,
		},
	}
}
