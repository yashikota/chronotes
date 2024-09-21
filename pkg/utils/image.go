package utils

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"

	"golang.org/x/image/draw"
)

func LoadImage(data []byte, fileType string) (image.Image, error) {
	// Readable image formats: jpeg, png
	var img image.Image
	var err error

	switch fileType {
	case "image/jpeg":
		img, err = jpeg.Decode(bytes.NewReader(data))
	case "image/png":
		img, err = png.Decode(bytes.NewReader(data))
	default:
		err = errors.New("unsupported file type")
	}

	if err != nil {
		return nil, errors.New("failed to decode the image: " + err.Error())
	}

	return img, nil
}

func Resize(data []byte, maxHeight int) []byte {
	// Aspect ratio preserving image resizing
	mimeType := http.DetectContentType(data)
	img, err := LoadImage(data, mimeType)
	if err != nil {
		return nil
	}

	// Calculate the new size
	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y

	if width > height {
		width = width * maxHeight / height
		height = maxHeight
	} else {
		height = height * maxHeight / width
		width = maxHeight
	}

	// Resize
	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.CatmullRom.Scale(dst, dst.Bounds(), img, img.Bounds(), draw.Over, nil)

	// Encode the image
	buf := new(bytes.Buffer)
	png.Encode(buf, dst)

	return buf.Bytes()
}
