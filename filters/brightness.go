package filters

import (
	"io"
	"sync"

	"github.com/julyskies/brille/v2/utilities"
)

func Brightness(file io.Reader, amount int) (io.Reader, string, error) {
	img, format, convertationError := utilities.DecodeSource(file)
	if convertationError != nil {
		return nil, "", convertationError
	}
	amount = utilities.Clamp(amount, 255, -255)
	pixLen := len(img.Pix)
	threads := utilities.GetThreads()
	pixPerThread := utilities.GetPixPerThread(pixLen, threads)
	var wg sync.WaitGroup
	processing := func(thread int) {
		defer wg.Done()
		startIndex := pixPerThread * thread
		endIndex := utilities.ClampMax(startIndex+pixPerThread, pixLen)
		for i := startIndex; i < endIndex; i += 4 {
			img.Pix[i] = uint8(utilities.Clamp(int(img.Pix[i])+amount, 255, 0))
			img.Pix[i+1] = uint8(utilities.Clamp(int(img.Pix[i+1])+amount, 255, 0))
			img.Pix[i+2] = uint8(utilities.Clamp(int(img.Pix[i+2])+amount, 255, 0))
		}
	}
	for t := 0; t < threads; t += 1 {
		wg.Add(1)
		go processing(t)
	}
	wg.Wait()
	return utilities.EncodeResult(img, format)
}
