package filters

import (
	"io"

	"github.com/julyskies/brille/constants"
	"github.com/julyskies/brille/utilities"
)

func Grayscale(file io.Reader, grayscaleType string) (io.Reader, string, error) {
	img, format, convertationError := utilities.DecodeSource(file)
	if convertationError != nil {
		return nil, "", convertationError
	}
	if grayscaleType != constants.GRAYSCALE_TYPE_AVERAGE &&
		grayscaleType != constants.GRAYSCALE_TYPE_LUMINANCE {
		grayscaleType = constants.GRAYSCALE_TYPE_AVERAGE
	}
	for i := 0; i < len(img.Pix); i += 4 {
		var channel uint8
		if grayscaleType == constants.GRAYSCALE_TYPE_AVERAGE {
			channel = uint8((int(img.Pix[i]) + int(img.Pix[i+1]) + int(img.Pix[i+2])) / 3)
		} else {
			channel = uint8(
				float64(img.Pix[i])*0.21 + float64(img.Pix[i+1])*0.72 + float64(img.Pix[i+2])*0.07,
			)
		}
		img.Pix[i], img.Pix[i+1], img.Pix[i+2] = channel, channel, channel
	}
	return utilities.EncodeResult(img, format)
}
