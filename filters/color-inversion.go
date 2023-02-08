package filters

import (
	"io"

	"github.com/julyskies/brille/v2/utilities"
)

func ColorInversion(file io.Reader) (io.Reader, string, error) {
	img, format, convertationError := utilities.DecodeSource(file)
	if convertationError != nil {
		return nil, "", convertationError
	}
	for i := 0; i < len(img.Pix); i += 4 {
		img.Pix[i] = 255 - img.Pix[i]
		img.Pix[i+1] = 255 - img.Pix[i+1]
		img.Pix[i+2] = 255 - img.Pix[i+2]
	}
	return utilities.EncodeResult(img, format)
}
