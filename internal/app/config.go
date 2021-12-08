package app

import (
	"fmt"

	"github.com/joeshaw/envdecode"
	_ "github.com/joho/godotenv/autoload"
)

type AppConfig struct {
	Server serverConfig
	Db     dbConfig
}

type dbConfig struct {
	Host     string `env:"DB_HOST,default=localhost"`
	Port     int    `env:"DB_PORT,default=9001"`
	Username string `env:"DB_USERNAME,default=manny"`
	Password string `env:"DB_PASSWORD,default=change-in-production"`
	DbName   string `env:"DB_NAME,default=brain"`
}

type serverConfig struct {
	Host string `env:"SERVER_HOST,default=localhost"`
	Port int    `env:"SERVER_PORT,default=3001"`
}

func LoadAppConfig() *AppConfig {
	var c AppConfig
	if err := envdecode.StrictDecode(&c); err != nil {
		fmt.Printf("Failed to decode: %s", err)
		panic(err)
	}
	return &c
}
