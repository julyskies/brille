package filters

import (
	"io"

	"github.com/julyskies/brille/v2/utilities"
)

func Contrast(file io.Reader, amount int) (io.Reader, string, error) {
	img, format, convertationError := utilities.DecodeSource(file)
	if convertationError != nil {
		return nil, "", convertationError
	}
	amount = utilities.MaxMin(amount, 255, -255)
	factor := float64(259*(amount+255)) / float64(255*(259-amount))
	for i := 0; i < len(img.Pix); i += 4 {
		img.Pix[i] = uint8(utilities.MaxMin(factor*(float64(img.Pix[i])-128)+128, 255, 0))
		img.Pix[i+1] = uint8(utilities.MaxMin(factor*(float64(img.Pix[i+1])-128)+128, 255, 0))
		img.Pix[i+2] = uint8(utilities.MaxMin(factor*(float64(img.Pix[i+2])-128)+128, 255, 0))
	}
	return utilities.EncodeResult(img, format)
}
