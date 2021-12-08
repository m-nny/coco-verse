package server

import (
	"github.com/gin-gonic/gin"
	"github.com/m-nny/coco-verse/internal/albums"
	"github.com/m-nny/coco-verse/internal/app"
)

func NewRouter(c *app.Context) *gin.Engine {

	router := gin.Default()

	albumsRepo := albums.NewRepository(c.Db)
	albumsController := albums.NewController(albumsRepo)

	albumsController.Bind(router)
	return router
}
