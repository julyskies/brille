package filters

import (
	"io"
	"sync"

	"github.com/julyskies/brille/v2/utilities"
)

func applySolarizeThreshold(subpixel, threshold uint8) uint8 {
	if subpixel < threshold {
		return 255 - subpixel
	}
	return subpixel
}

func Solarize(file io.Reader, threshold uint8) (io.Reader, string, error) {
	img, format, convertationError := utilities.DecodeSource(file)
	if convertationError != nil {
		return nil, "", convertationError
	}
	pixLen := len(img.Pix)
	threads := utilities.GetThreads()
	pixPerThread := utilities.GetPixPerThread(pixLen, threads)
	var wg sync.WaitGroup
	processing := func(thread int) {
		defer wg.Done()
		startIndex := pixPerThread * thread
		endIndex := utilities.ClampMax(startIndex+pixPerThread, pixLen)
		for i := startIndex; i < endIndex; i += 4 {
			img.Pix[i] = applySolarizeThreshold(img.Pix[i], threshold)
			img.Pix[i+1] = applySolarizeThreshold(img.Pix[i+1], threshold)
			img.Pix[i+2] = applySolarizeThreshold(img.Pix[i+2], threshold)
		}
	}
	for t := 0; t < threads; t += 1 {
		wg.Add(1)
		go processing(t)
	}
	wg.Wait()
	return utilities.EncodeResult(img, format)
}
