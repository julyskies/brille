package utilities

import (
	"image"
	"image/draw"
	"io"
)

func DecodeSource(file io.Reader) (*image.RGBA, string, error) {
	content, format, decodingError := image.Decode(file)
	if decodingError != nil {
		return nil, "", decodingError
	}
	rect := content.Bounds()
	img := image.NewRGBA(rect)
	draw.Draw(img, img.Bounds(), content, rect.Min, draw.Src)
	return img, format, nil
}
