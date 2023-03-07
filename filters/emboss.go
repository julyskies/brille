package filters

import (
	"io"
	"math"
	"sync"

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
	pixLen := len(img.Pix)
	result := make([]uint8, pixLen)
	threads := utilities.GetThreads()
	pixPerThread := utilities.GetPixPerThread(pixLen, threads)
	var wg sync.WaitGroup
	processing := func(thread int) {
		defer wg.Done()
		startIndex := pixPerThread * thread
		endIndex := utilities.ClampMax(startIndex+pixPerThread, pixLen)
		for i := startIndex; i < endIndex; i += 4 {
			x, y := utilities.GetCoordinates(i/4, width)
			gradientX := 0
			gradientY := 0
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
			channel := uint8(
				255 - utilities.ClampMax(
					math.Sqrt(float64(gradientX*gradientX+gradientY*gradientY)),
					255,
				),
			)
			result[i], result[i+1], result[i+2], result[i+3] = channel, channel, channel, img.Pix[i+3]
		}
	}
	for t := 0; t < threads; t += 1 {
		wg.Add(1)
		go processing(t)
	}
	wg.Wait()
	img.Pix = result
	return utilities.EncodeResult(img, format)
}
