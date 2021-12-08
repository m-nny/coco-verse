package app

import (
	"github.com/m-nny/coco-verse/internal/albums"
	"github.com/m-nny/coco-verse/internal/config"
	g "github.com/m-nny/coco-verse/internal/gorm"
	"gorm.io/gorm"
)

type Context struct {
	Config *config.AppConfig
	Db     *gorm.DB
}

func NewContext() *Context {
	appConfig := config.Load()
	db := g.Connect(appConfig)

	db.AutoMigrate(&albums.Album{})

	return &Context{Config: appConfig, Db: db}
}
