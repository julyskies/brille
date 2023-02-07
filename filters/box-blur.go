package filters

import (
	"io"

	"github.com/julyskies/brille/utilities"
)

func BoxBlur(file io.Reader, radius uint) (io.Reader, string, error) {
	img, format, convertationError := utilities.DecodeSource(file)
	if convertationError != nil {
		return nil, "", convertationError
	}
	radiusInt := int(radius)
	width, height := img.Rect.Max.X, img.Rect.Max.Y
	for i := 0; i < len(img.Pix); i += 4 {
		x, y := utilities.GetCoordinates(i/4, width)
		sumR, sumG, sumB, pixelCount := 0, 0, 0, 0
		x2s, x2e := utilities.GetAperture(x, width, -radiusInt, radiusInt)
		y2s, y2e := utilities.GetAperture(y, height, -radiusInt, radiusInt)
		for x2 := x2s; x2 < x2e; x2 += 1 {
			for y2 := y2s; y2 < y2e; y2 += 1 {
				px := utilities.GetPixel(x2, y2, width)
				sumR += int(img.Pix[px])
				sumG += int(img.Pix[px+1])
				sumB += int(img.Pix[px+2])
				pixelCount += 1
			}
		}
		img.Pix[i] = uint8(sumR / pixelCount)
		img.Pix[i+1] = uint8(sumG / pixelCount)
		img.Pix[i+2] = uint8(sumB / pixelCount)
	}
	return utilities.EncodeResult(img, format)
}
