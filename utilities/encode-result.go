package utilities

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strconv"

	"github.com/julyskies/brille/v2/constants"
)

func getJPEGQuality() int {
	jpegQualityENV := os.Getenv(constants.ENV_JPEG_QUALITY)
	jpegQuality := constants.ENV_JPEG_QUALITY_DEFAULT
	if jpegQualityENV != "" {
		parsed, parsingError := strconv.Atoi(jpegQualityENV)
		if parsingError == nil {
			jpegQuality = MaxMin(parsed, 100, 0)
		}
	}
	return jpegQuality
}

func EncodeResult(img *image.RGBA, format string) (io.Reader, string, error) {
	var buffer bytes.Buffer
	writer := io.Writer(&buffer)
	if format == "png" {
		encodingError := png.Encode(writer, img.SubImage(img.Rect))
		if encodingError != nil {
			return nil, "", encodingError
		}
	} else {
		jpegQuality := getJPEGQuality()
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
