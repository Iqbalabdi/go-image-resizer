package utils

import (
	"bytes"
	"encoding/base64"
	"image/jpeg"
	"image/png"
	"log"

	u "github.com/Iqbalabdi/go-image-resizer/upsert"
)

func toBase64(b[] byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func imgToBytes(img u.ImageData) []byte {
	var buf bytes.Buffer
	switch img.ImgType {
	case "jpeg":
		err := jpeg.Encode(&buf, img.Data, nil)
		if err != nil {
			log.Println(err.Error())
		}
	case "png":
		err := png.Encode(&buf, img.Data)
		if err != nil {
			log.Println(err.Error())
		}
	}
	return buf.Bytes()
}

func ConvertToBase64(img u.ImageData) string {
	imgByte := imgToBytes(img)
	bs64strings := "data:image/"+ img.ImgType + ";base64," + toBase64(imgByte)
	return bs64strings
}