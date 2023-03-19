package filters

import (
	"io"
	"math"
	"sync"

	"github.com/julyskies/brille/v2/utilities"
)

const K float64 = 6

func createKernel(sigma float64) []float64 {
	dim := math.Max(3.0, K*sigma)
	sqrtSigmaPi2 := math.Sqrt(math.Pi*2.0) * sigma
	s2 := 2.0 * sigma * sigma
	sum := 0.0
	kDim := dim
	if int(kDim)%2 == 0 {
		kDim = dim - 1
	}
	kernel := make([]float64, int(kDim))
	half := int(len(kernel) / 2)
	i := -half
	for j := 0; j < len(kernel); j += 1 {
		kernel[j] = math.Exp(-float64(i*i)/(s2)) / sqrtSigmaPi2
		sum += kernel[j]
		i += 1
	}
	for k := 0; k < int(kDim); k += 1 {
		kernel[k] /= sum
	}
	return kernel
}

func GaussianBlur(file io.Reader, sigma float64) (io.Reader, string, error) {
	img, format, convertationError := utilities.DecodeSource(file)
	if convertationError != nil {
		return nil, "", convertationError
	}
	kernel := createKernel(sigma)
	pixLen := len(img.Pix)
	width, height := img.Rect.Max.X, img.Rect.Max.Y
	temp := make([]uint8, pixLen)
	threads := utilities.GetThreads()
	pixPerThread := utilities.GetPixPerThread(pixLen, threads)
	var wg sync.WaitGroup

	processing := func(start int, direction string) {
		defer wg.Done()
		end := utilities.ClampMax(start+pixPerThread, pixLen)
		for i := start; i < end; i += 4 {
			x, y := utilities.GetCoordinates(i/4, width)
			sumR := 0.0
			sumG := 0.0
			sumB := 0.0
			for k := 0; k < len(kernel); k += 1 {
				var px int
				if direction == "horizontal" {
					px = utilities.GetPixel(
						utilities.Clamp(x-(len(kernel)/2-k), width-1, 0),
						y,
						width,
					)
					sumR += float64(img.Pix[px]) * kernel[k]
					sumG += float64(img.Pix[px+1]) * kernel[k]
					sumB += float64(img.Pix[px+2]) * kernel[k]
				} else {
					px = utilities.GetPixel(
						x,
						utilities.Clamp(y-(len(kernel)/2-k), height-1, 0),
						width,
					)
					sumR += float64(temp[px]) * kernel[k]
					sumG += float64(temp[px+1]) * kernel[k]
					sumB += float64(temp[px+2]) * kernel[k]
				}
			}
			if direction == "horizontal" {
				temp[i] = uint8(utilities.Clamp(sumR, 255, 0))
				temp[i+1] = uint8(utilities.Clamp(sumG, 255, 0))
				temp[i+2] = uint8(utilities.Clamp(sumB, 255, 0))
			} else {
				img.Pix[i] = uint8(utilities.Clamp(sumR, 255, 0))
				img.Pix[i+1] = uint8(utilities.Clamp(sumG, 255, 0))
				img.Pix[i+2] = uint8(utilities.Clamp(sumB, 255, 0))
			}
		}
	}

	// horizontal
	for t := 0; t < threads; t += 1 {
		wg.Add(1)
		go processing(pixPerThread*t, "horizontal")
	}
	wg.Wait()

	// vertical
	for t := 0; t < threads; t += 1 {
		wg.Add(1)
		go processing(pixPerThread*t, "vertical")
	}
	wg.Wait()

	return utilities.EncodeResult(img, format)
}
