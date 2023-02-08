package filters

import (
	"io"
	"math"

	"github.com/julyskies/brille/v2/utilities"
)

var embossHorizontal = [3][3]int{
	{0, 0, 0},
	{1, 0, -1},
	{0, 0, 0},
}

var embossVertical = [3][3]int{
	{0, 1, 0},
	{0, 0, 0},
	{0, -1, 0},
}

func Emboss(file io.Reader) (io.Reader, string, error) {
	img, format, convertationError := utilities.DecodeSource(file)
	if convertationError != nil {
		return nil, "", convertationError
	}
	width, height := img.Rect.Max.X, img.Rect.Max.Y
	for i := 0; i < len(img.Pix); i += 4 {
		x, y := utilities.GetCoordinates(i/4, width)
		gradientX, gradientY := 0, 0
		for m := 0; m < 3; m += 1 {
			for n := 0; n < 3; n += 1 {
				k := utilities.GradientPoint(x, m, width)
				l := utilities.GradientPoint(y, n, height)
				px := utilities.GetPixel(x+k, y+l, width)
				average := (int(img.Pix[px]) + int(img.Pix[px+1]) + int(img.Pix[px+2])) / 3
				gradientX += average * embossHorizontal[m][n]
				gradientY += average * embossVertical[m][n]
			}
		}
		colorCode := uint8(
			255 - utilities.MaxMin(
				math.Sqrt(float64(gradientX*gradientX+gradientY*gradientY)),
				255,
				0,
			),
		)
		img.Pix[i], img.Pix[i+1], img.Pix[i+2] = colorCode, colorCode, colorCode
	}
	return utilities.EncodeResult(img, format)
}
