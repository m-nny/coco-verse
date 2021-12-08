package server

import (
	"github.com/gin-gonic/gin"
	"github.com/m-nny/coco-verse/internal/albums"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	albumsRepo := albums.NewRepository()
	albumsController := albums.NewController(albumsRepo)

	albumsController.Bind(router)
	return router
}
