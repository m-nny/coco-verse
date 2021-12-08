package config

import (
	"fmt"

	"github.com/joeshaw/envdecode"
	_ "github.com/joho/godotenv/autoload"
)

type AppConfig struct {
	Server serverConfig
	Db     dbConfig
}

func Load() *AppConfig {
	var c AppConfig
	if err := envdecode.StrictDecode(&c); err != nil {
		fmt.Printf("Failed to decode: %s", err)
		panic(err)
	}
	return &c
}
