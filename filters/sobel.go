package filters

import (
	"io"
	"math"
	"sync"

	"github.com/julyskies/brille/v2/utilities"
)

var sobelHorizontal = [3][3]int{
	{-1, 0, 1},
	{-2, 0, 2},
	{-1, 0, 1},
}

var sobelVertical = [3][3]int{
	{1, 2, 1},
	{0, 0, 0},
	{-1, -2, -1},
}

func Sobel(file io.Reader) (io.Reader, string, error) {
	img, format, convertationError := utilities.DecodeSource(file)
	if convertationError != nil {
		return nil, "", convertationError
	}
	pixLen := len(img.Pix)
	threads := utilities.GetThreads()
	pixPerThread := utilities.GetPixPerThread(pixLen, threads)
	width, height := img.Rect.Max.X, img.Rect.Max.Y
	result := make([]uint8, pixLen)
	var wg sync.WaitGroup
	processing := func(thread int) {
		defer wg.Done()
		startIndex := pixPerThread * thread
		endIndex := utilities.ClampMax(startIndex+pixPerThread, pixLen)
		for i := startIndex; i < endIndex; i += 4 {
			x, y := utilities.GetCoordinates(i/4, width)
			gradientX, gradientY := 0, 0
			for m := 0; m < 3; m += 1 {
				for n := 0; n < 3; n += 1 {
					px := utilities.GetPixel(
						utilities.Clamp(x-(len(laplacianKernel)/2-m), width-1, 0),
						utilities.Clamp(y-(len(laplacianKernel)/2-n), height-1, 0),
						width,
					)
					average := (int(img.Pix[px]) + int(img.Pix[px+1]) + int(img.Pix[px+2])) / 3
					gradientX += average * sobelHorizontal[m][n]
					gradientY += average * sobelVertical[m][n]
				}
			}
			channel := uint8(
				255 - utilities.Clamp(
					math.Sqrt(float64(gradientX*gradientX+gradientY*gradientY)),
					255,
					0,
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
