package utilities

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
)

func EncodeGridResult(result [][]color.Color, format string) (io.Reader, string, error) {
	width, height := len(result), len(result[0])
	nrgba := image.NewNRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			nrgba.Set(x, y, result[x][y])
		}
	}
	var buffer bytes.Buffer
	writer := io.Writer(&buffer)
	if format == "png" {
		encodingError := png.Encode(writer, nrgba.SubImage(nrgba.Rect))
		if encodingError != nil {
			return nil, "", encodingError
		}
	} else {
		jpegQuality := getJPEGQuality()
		encodingError := jpeg.Encode(
			writer,
			nrgba.SubImage(nrgba.Rect),
			&jpeg.Options{
				Quality: jpegQuality,
			},
		)
		if encodingError != nil {
			return nil, "", encodingError
		}
	}
	return bytes.NewReader(buffer.Bytes()), format, nil
}
