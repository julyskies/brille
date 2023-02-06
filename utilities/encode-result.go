package utilities

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strconv"
)

func EncodeResult(img *image.RGBA, format string) (io.Reader, string, error) {
	jpegQualityENV := os.Getenv("BRILLE_JPEG_QUALITY")
	jpegQuality := 100
	if jpegQualityENV != "" {
		parsed, parsingError := strconv.Atoi(jpegQualityENV)
		if parsingError == nil {
			jpegQuality = MaxMin(parsed, 100, 0)
		}
	}
	var buffer bytes.Buffer
	writer := io.Writer(&buffer)
	if format == "png" {
		encodingError := png.Encode(writer, img.SubImage(img.Rect))
		if encodingError != nil {
			return nil, "", encodingError
		}
	} else {
		encodingError := jpeg.Encode(
			writer,
			img.SubImage(img.Rect),
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
