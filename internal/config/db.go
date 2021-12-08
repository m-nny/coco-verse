package config

import "fmt"

type dbConfig struct {
	Host     string `env:"DB_HOST,default=localhost"`
	Port     int    `env:"DB_PORT,default=9001"`
	Username string `env:"DB_USERNAME,default=manny"`
	Password string `env:"DB_PASSWORD,default=change-in-production"`
	DbName   string `env:"DB_NAME,default=brain"`
}

func (dbConfig *dbConfig) Dsn() string {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Almaty",
		dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.DbName, dbConfig.Port)
	return dsn
}
