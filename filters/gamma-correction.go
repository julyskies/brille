package filters

import (
	"io"
	"math"
	"sync"

	"github.com/julyskies/brille/v2/utilities"
)

func GammaCorrection(file io.Reader, amount float64) (io.Reader, string, error) {
	img, format, convertationError := utilities.DecodeSource(file)
	if convertationError != nil {
		return nil, "", convertationError
	}
	amount = utilities.Clamp(amount, 3.99, 0)
	power := 1 / amount
	pixLen := len(img.Pix)
	threads := utilities.GetThreads()
	pixPerThread := utilities.GetPixPerThread(pixLen, threads)
	var wg sync.WaitGroup
	processing := func(thread int) {
		defer wg.Done()
		startIndex := pixPerThread * thread
		endIndex := utilities.ClampMax(startIndex+pixPerThread, pixLen)
		for i := startIndex; i < endIndex; i += 4 {
			img.Pix[i] = uint8(255 * math.Pow(float64(img.Pix[i])/255, power))
			img.Pix[i+1] = uint8(255 * math.Pow(float64(img.Pix[i+1])/255, power))
			img.Pix[i+2] = uint8(255 * math.Pow(float64(img.Pix[i+2])/255, power))
		}
	}
	for t := 0; t < threads; t += 1 {
		wg.Add(1)
		go processing(t)
	}
	wg.Wait()
	return utilities.EncodeResult(img, format)
}
