package albums

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlbumController struct {
	repo *AlbumRepository
}

func (controller *AlbumController) getAll(c *gin.Context) {
	items := controller.repo.FindAll()
	c.IndentedJSON(http.StatusOK, items)
}
func (controller *AlbumController) getById(c *gin.Context) {
	id := c.Param("id")

	item := controller.repo.FindById(id)
	if item == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	} else {
		c.IndentedJSON(http.StatusOK, item)
	}
}
func (controller *AlbumController) postOne(c *gin.Context) {
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	controller.repo.AddOne(&newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func NewController(repo *AlbumRepository) *AlbumController {
	return &AlbumController{
		repo: repo,
	}
}

func (controller *AlbumController) Bind(router *gin.Engine) {
	group := router.Group("/albums")
	{
		group.GET("/", controller.getAll)
		group.GET("/:id", controller.getById)
		group.POST("/", controller.postOne)
	}
}
