package utils

import (
	"image"

	"github.com/disintegration/imaging"
)

func ResizeImage(img image.Image) image.Image {
	resizedImage := imaging.Resize(img, 250, 250, imaging.Lanczos)
	return resizedImage
}