package filters

import (
	"io"
	"math"
	"sync"

	"github.com/julyskies/brille/v2/utilities"
)

const DEG float64 = math.Pi / 180

func HueRotate(file io.Reader, angle int) (io.Reader, string, error) {
	img, format, convertationError := utilities.DecodeSource(file)
	if convertationError != nil {
		return nil, "", convertationError
	}
	cos := math.Cos(float64(angle) * DEG)
	sin := math.Sin(float64(angle) * DEG)
	matrix := [3]float64{
		cos + (1-cos)/3,
		(1-cos)/3 - math.Sqrt(float64(1)/3)*sin,
		(1-cos)/3 + math.Sqrt(float64(1)/3)*sin,
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
			rr := utilities.Clamp(
				float64(r)*matrix[0]+float64(g)*matrix[1]+float64(b)*matrix[2],
				255,
				0,
			)
			rg := utilities.Clamp(
				float64(r)*matrix[2]+float64(g)*matrix[0]+float64(b)*matrix[1],
				255,
				0,
			)
			rb := utilities.Clamp(
				float64(r)*matrix[1]+float64(g)*matrix[2]+float64(b)*matrix[0],
				255,
				0,
			)
			img.Pix[i] = uint8(rr)
			img.Pix[i+1] = uint8(rg)
			img.Pix[i+2] = uint8(rb)
		}
	}
	for t := 0; t < threads; t += 1 {
		wg.Add(1)
		go processing(t)
	}
	wg.Wait()
	return utilities.EncodeResult(img, format)
}
