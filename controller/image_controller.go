package controller

import (
	"fmt"
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
	// width := form.Value["width"][0]
	// height := form.Value["height"][0]
	fmt.Println(files)
	resizedImage, err := imc.imageService.Resize(files)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"img": resizedImage,
	})
}

func (imc *ImageController) HandleConvertToJPEG(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	convertImage, err := imc.imageService.ConvertToJPEG(files)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"img": convertImage,
	})
}