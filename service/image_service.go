package service

import (
	"bytes"
	"image"
	"image/png"
	"io"
	"log"
	"mime/multipart"
	"github.com/Iqbalabdi/go-image-resizer/utils"
	u "github.com/Iqbalabdi/go-image-resizer/upsert"
)

type ImageService struct{}

func NewImageService() *ImageService {
	return &ImageService{}
}

func (ims *ImageService) Resize(files []*multipart.FileHeader)([]string, error) {
	var imgs []u.ImageData
	var base64Strings []string
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

		img, imgType, err := image.Decode(bytes.NewReader(flRead))
		if err != nil {
			log.Println("Couldn't decode image")
			return base64Strings, err
		}
		resizedImage := utils.ResizeImage(img)
		imgs = append(imgs, u.ImageData{
			Data: resizedImage,
			ImgType: imgType,
		})
	}
	for _, img := range imgs {
		base64Strings = append(base64Strings, utils.ConvertToBase64(img))
	}
	return base64Strings, nil
}

func (ims *ImageService) ConvertToJPEG(files []*multipart.FileHeader)([]string, error) {
	var imgs []u.ImageData
	var base64Strings []string
	for _, file := range files {
		fl, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer fl.Close()
		flRead, err := io.ReadAll(fl)
		if err != nil {
			log.Println("Couldn't open file")
			return nil, err
		}
		img, err := png.Decode(bytes.NewReader(flRead))
		if err != nil {
			log.Println("Unable to decode PNG")
			return nil, err
		}
		imgs = append(imgs, u.ImageData{
			Data: img,
			ImgType: "jpeg",
		})
	}
	for _, img := range imgs {
		base64Strings = append(base64Strings, utils.ConvertToBase64(img))
	}
	return base64Strings, nil
}