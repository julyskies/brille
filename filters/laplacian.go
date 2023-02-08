package filters

import (
	"io"

	"github.com/julyskies/brille/v2/utilities"
)

var laplacianKernel = [3][3]int{
	{-1, -1, -1},
	{-1, 8, -1},
	{-1, -1, -1},
}

func Laplacian(file io.Reader) (io.Reader, string, error) {
	img, format, convertationError := utilities.DecodeSource(file)
	if convertationError != nil {
		return nil, "", convertationError
	}
	width, height := img.Rect.Max.X, img.Rect.Max.Y
	for i := 0; i < len(img.Pix); i += 4 {
		averageSum := 0
		x, y := utilities.GetCoordinates(i/4, width)
		for m := 0; m < 3; m += 1 {
			for n := 0; n < 3; n += 1 {
				k := utilities.GradientPoint(x, m, width)
				l := utilities.GradientPoint(y, n, height)
				px := utilities.GetPixel(x+k, y+l, width)
				average := (int(img.Pix[px]) + int(img.Pix[px+1]) + int(img.Pix[px+2])) / 3
				averageSum += average * laplacianKernel[m][n]
			}
		}
		channel := 255 - uint8(utilities.MaxMin(averageSum, 255, 0))
		img.Pix[i], img.Pix[i+1], img.Pix[i+2] = channel, channel, channel
	}
	return utilities.EncodeResult(img, format)
}
