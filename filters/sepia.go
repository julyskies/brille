package filters

import (
	"io"
	"sync"

	"github.com/julyskies/brille/v2/utilities"
)

func Sepia(file io.Reader) (io.Reader, string, error) {
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
			r, g, b := img.Pix[i], img.Pix[i+1], img.Pix[i+2]
			dR := utilities.ClampMax(0.393*float64(r)+0.769*float64(g)+0.189*float64(b), 255.0)
			dG := utilities.ClampMax(0.349*float64(r)+0.686*float64(g)+0.168*float64(b), 255.0)
			dB := utilities.ClampMax(0.272*float64(r)+0.534*float64(g)+0.131*float64(b), 255.0)
			img.Pix[i], img.Pix[i+1], img.Pix[i+2] = uint8(dR), uint8(dG), uint8(dB)
		}
	}
	for t := 0; t < threads; t += 1 {
		wg.Add(1)
		go processing(t)
	}
	wg.Wait()
	return utilities.EncodeResult(img, format)
}
