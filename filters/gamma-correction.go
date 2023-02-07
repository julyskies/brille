package filters

import (
	"io"
	"math"

	"github.com/julyskies/brille/utilities"
)

func GammaCorrection(file io.Reader, amount float64) (io.Reader, string, error) {
	img, format, convertationError := utilities.DecodeSource(file)
	if convertationError != nil {
		return nil, "", convertationError
	}
	amount = utilities.MaxMin(amount, 3.99, 0)
	power := 1 / amount
	for i := 0; i < len(img.Pix); i += 4 {
		img.Pix[i] = uint8(255 * math.Pow(float64(img.Pix[i])/255, power))
		img.Pix[i+1] = uint8(255 * math.Pow(float64(img.Pix[i+1])/255, power))
		img.Pix[i+2] = uint8(255 * math.Pow(float64(img.Pix[i+2])/255, power))
	}
	return utilities.EncodeResult(img, format)
}
