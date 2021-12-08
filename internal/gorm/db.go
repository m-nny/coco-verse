package gorm

import (
	"fmt"

	"github.com/m-nny/coco-verse/internal/app"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(c *app.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Almaty", c.Db.Host, c.Db.Username, c.Db.Password, c.Db.DbName, c.Db.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connecto to database")
	}

	return db
}
