package albums

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AlbumController struct {
	repo *AlbumRepository
}

func (controller *AlbumController) getAll(c *gin.Context) {
	items, err := controller.repo.FindAll()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Error", "__internal": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, items)
	}
}
func (controller *AlbumController) getById(c *gin.Context) {
	id := c.Param("id")

	item, err := controller.repo.FindById(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Error", "__internal": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, item)
	}
}
func (controller *AlbumController) postOne(c *gin.Context) {
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Internal Error", "__internal": err.Error()})
		return
	}
	err := controller.repo.AddOne(&newAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Error", "__internal": err.Error()})
	} else {
		c.IndentedJSON(http.StatusCreated, newAlbum)
	}
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
