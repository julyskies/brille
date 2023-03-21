package filters

import (
	"io"
	"sync"

	"github.com/julyskies/brille/v2/utilities"
)

var sharpenKernel = [3][3]int{
	{-1, -1, -1},
	{-1, 9, -1},
	{-1, -1, -1},
}

func Sharpen(file io.Reader, amount uint) (io.Reader, string, error) {
	img, format, convertationError := utilities.DecodeSource(file)
	if convertationError != nil {
		return nil, "", convertationError
	}
	mix := float64(utilities.Clamp(amount, 100, 0)) / 100
	pixLen := len(img.Pix)
	threads := utilities.GetThreads()
	pixPerThread := utilities.GetPixPerThread(pixLen, threads)
	result := make([]uint8, pixLen)
	width, height := img.Rect.Max.X, img.Rect.Max.Y
	var wg sync.WaitGroup
	processing := func(thread int) {
		defer wg.Done()
		startIndex := pixPerThread * thread
		endIndex := utilities.ClampMax(startIndex+pixPerThread, pixLen)
		for i := startIndex; i < endIndex; i += 4 {
			dR, dG, dB := 0, 0, 0
			x, y := utilities.GetCoordinates(i/4, width)
			for m := 0; m < 3; m += 1 {
				for n := 0; n < 3; n += 1 {
					k := utilities.GradientPoint(x, m, width)
					l := utilities.GradientPoint(y, n, height)
					px := utilities.GetPixel(x+k, y+l, width)
					dR += int(img.Pix[px]) * sharpenKernel[m][n]
					dG += int(img.Pix[px+1]) * sharpenKernel[m][n]
					dB += int(img.Pix[px+2]) * sharpenKernel[m][n]
				}
			}
			result[i] = uint8(
				utilities.Clamp(float64(dR)*mix+float64(img.Pix[i])*(1-mix), 255, 0),
			)
			result[i+1] = uint8(
				utilities.Clamp(float64(dG)*mix+float64(img.Pix[i+1])*(1-mix), 255, 0),
			)
			result[i+2] = uint8(
				utilities.Clamp(float64(dB)*mix+float64(img.Pix[i+2])*(1-mix), 255, 0),
			)
			result[i+3] = img.Pix[i+3]
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
