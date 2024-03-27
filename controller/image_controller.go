package controller

import (
	"net/http"

	"github.com/Iqbalabdi/go-image-resizer/service"
	"github.com/gin-gonic/gin"
)

type ImageController struct {
	imageService *service.ImageService
}

func NewImageController(service *service.ImageService) *ImageController {
	return &ImageController{
		imageService: service,
	}
}

func (imc *ImageController) HandleResize(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	resizedImage, err := imc.imageService.Resize(files)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Couldn't resize image",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"img": resizedImage,
	})
}