package filters

import (
	"io"

	"github.com/julyskies/brille/v2/utilities"
)

func applySolarizeThreshold(subpixel, threshold uint8) uint8 {
	if subpixel < threshold {
		return 255 - subpixel
	}
	return subpixel
}

func Solarize(file io.Reader, threshold uint8) (io.Reader, string, error) {
	img, format, convertationError := utilities.DecodeSource(file)
	if convertationError != nil {
		return nil, "", convertationError
	}
	for i := 0; i < len(img.Pix); i += 4 {
		img.Pix[i] = applySolarizeThreshold(img.Pix[i], threshold)
		img.Pix[i+1] = applySolarizeThreshold(img.Pix[i+1], threshold)
		img.Pix[i+2] = applySolarizeThreshold(img.Pix[i+2], threshold)
	}
	return utilities.EncodeResult(img, format)
}
