package utils

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"log"
)

func toBase64(b[] byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func imgToBytes(img image.Image) []byte {
	var buf bytes.Buffer
	err := jpeg.Encode(&buf, img, nil)
	if err != nil {
		log.Println("Couldn't encode image")
	}
	return buf.Bytes()
}

func ConvertToBase64(img image.Image) string {
	imgByte := imgToBytes(img)
	bs64strings := "data:image/jpeg;base64" + toBase64(imgByte)
	return bs64strings
}