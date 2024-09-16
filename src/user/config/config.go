package config

import (
	"log"
	"os"

	"github.com/benosborntech/feedme/updates/consts"
)

type Config struct {
	Port     string
	MySQLDSN string
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

	return &Config{
		Port:     port,
		MySQLDSN: mySQLDSN,
	}
}
