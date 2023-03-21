package filters

import (
	"io"
	"sync"

	"github.com/julyskies/brille/v2/utilities"
)

func BoxBlur(file io.Reader, radius uint) (io.Reader, string, error) {
	img, format, convertationError := utilities.DecodeSource(file)
	if convertationError != nil {
		return nil, "", convertationError
	}
	radiusInt := int(radius)
	width, height := img.Rect.Max.X, img.Rect.Max.Y
	pixLen := len(img.Pix)
	threads := utilities.GetThreads()
	pixPerThread := utilities.GetPixPerThread(pixLen, threads)
	result := make([]uint8, pixLen)
	var wg sync.WaitGroup
	processing := func(thread int) {
		defer wg.Done()
		startIndex := pixPerThread * thread
		endIndex := utilities.ClampMax(startIndex+pixPerThread, pixLen)
		for i := startIndex; i < endIndex; i += 4 {
			x, y := utilities.GetCoordinates(i/4, width)
			dR, dG, dB := 0, 0, 0
			pixelCount := 0
			x2s, x2e := utilities.GetAperture(x, width, -radiusInt, radiusInt)
			y2s, y2e := utilities.GetAperture(y, height, -radiusInt, radiusInt)
			for x2 := x2s; x2 < x2e; x2 += 1 {
				for y2 := y2s; y2 < y2e; y2 += 1 {
					px := utilities.GetPixel(x2, y2, width)
					dR += int(img.Pix[px])
					dG += int(img.Pix[px+1])
					dB += int(img.Pix[px+2])
					pixelCount += 1
				}
			}
			result[i] = uint8(dR / pixelCount)
			result[i+1] = uint8(dG / pixelCount)
			result[i+2] = uint8(dB / pixelCount)
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
