package config

import "fmt"

type serverConfig struct {
	Host string `env:"SERVER_HOST,default=localhost"`
	Port int    `env:"SERVER_PORT,default=3001"`
}

func (serverConfig *serverConfig) GetBindAddress() string {
	return fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)
}
