package service

import (
	"bytes"
	"image"
	"image/jpeg"
	"io"
	"log"
	"mime/multipart"

	"github.com/Iqbalabdi/go-image-resizer/utils"
)

type ImageService struct{}

func NewImageService() *ImageService {
	return &ImageService{}
}

func (ims *ImageService) Resize(files []*multipart.FileHeader)([]string, error) {
	var imgs[] image.Image
	var base64Strings[] string
	for _, file := range files {
		fl, err := file.Open()
		if err != nil {
			return base64Strings, err
		}
		defer fl.Close()
		flRead, err := io.ReadAll(fl)
		if err != nil {
			log.Println("Couldn't open file")
			return base64Strings, err
		}
		m, err := jpeg.Decode(bytes.NewReader(flRead))
		if err != nil {
			log.Println("Couldn't decode image")
			return base64Strings, err
		}
		resizedImage := utils.ResizeImage(m)
		imgs = append(imgs, resizedImage)
	}
	for _, resizedImg := range imgs {
		base64Strings = append(base64Strings, utils.ConvertToBase64(resizedImg))
	}
	return base64Strings, nil
}
