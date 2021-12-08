package gorm

import (
	"github.com/m-nny/coco-verse/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(c *config.AppConfig) *gorm.DB {
	dsn := c.Db.Dsn()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connecto to database")
	}

	return db
}
