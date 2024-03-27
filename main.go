package main

import (
	"github.com/Iqbalabdi/go-image-resizer/service"
	"github.com/Iqbalabdi/go-image-resizer/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	imageService := service.NewImageService()
	imageController := controller.NewImageController(imageService)
	r.GET("/resize", imageController.HandleResize)
	r.Run("127.0.0.1:8080")
}