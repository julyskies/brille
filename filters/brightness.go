package filters

import (
	"io"

	"github.com/julyskies/brille/utilities"
)

func Brightness(file io.Reader, amount int) (io.Reader, string, error) {
	img, format, convertationError := utilities.DecodeSource(file)
	if convertationError != nil {
		return nil, "", convertationError
	}
	amount = utilities.MaxMin(amount, 255, -255)
	for i := 0; i < len(img.Pix); i += 4 {
		img.Pix[i] = uint8(utilities.MaxMin(int(img.Pix[i])+amount, 255, 0))
		img.Pix[i+1] = uint8(utilities.MaxMin(int(img.Pix[i+1])+amount, 255, 0))
		img.Pix[i+2] = uint8(utilities.MaxMin(int(img.Pix[i+2])+amount, 255, 0))
	}
	return utilities.EncodeResult(img, format)
}
