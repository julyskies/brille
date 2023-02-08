package filters

import (
	"io"

	"github.com/julyskies/brille/v2/utilities"
)

func Sepia(file io.Reader) (io.Reader, string, error) {
	img, format, convertationError := utilities.DecodeSource(file)
	if convertationError != nil {
		return nil, "", convertationError
	}
	for i := 0; i < len(img.Pix); i += 4 {
		r, g, b := img.Pix[i], img.Pix[i+1], img.Pix[i+2]
		sr := utilities.MaxMin(0.393*float64(r)+0.769*float64(g)+0.189*float64(b), 255.0, 0.0)
		sg := utilities.MaxMin(0.349*float64(r)+0.686*float64(g)+0.168*float64(b), 255.0, 0.0)
		sb := utilities.MaxMin(0.272*float64(r)+0.534*float64(g)+0.131*float64(b), 255.0, 0.0)
		img.Pix[i], img.Pix[i+1], img.Pix[i+2] = uint8(sr), uint8(sg), uint8(sb)
	}
	return utilities.EncodeResult(img, format)
}
