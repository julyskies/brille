package filters

import (
	"io"

	"github.com/julyskies/brille/v2/utilities"
)

func Binary(file io.Reader, threshold uint8) (io.Reader, string, error) {
	img, format, convertationError := utilities.DecodeSource(file)
	if convertationError != nil {
		return nil, "", convertationError
	}
	for i := 0; i < len(img.Pix); i += 4 {
		average := uint8((int(img.Pix[i]) + int(img.Pix[i+1]) + int(img.Pix[i+2])) / 3)
		channel := uint8(255)
		if average < threshold {
			channel = 0
		}
		img.Pix[i], img.Pix[i+1], img.Pix[i+2] = channel, channel, channel
	}
	return utilities.EncodeResult(img, format)
}
