package filters

import (
	"io"
	"sync"

	"github.com/julyskies/brille/v2/constants"
	"github.com/julyskies/brille/v2/utilities"
)

func Flip(file io.Reader, direction string) (io.Reader, string, error) {
	img, format, convertationError := utilities.DecodeSource(file)
	if convertationError != nil {
		return nil, "", convertationError
	}
	if direction != constants.FLIP_DIRECTION_HORIZONTAL &&
		direction != constants.FLIP_DIRECTION_VERTICAL {
		direction = constants.FLIP_DIRECTION_HORIZONTAL
	}
	width, height := img.Rect.Max.X, img.Rect.Max.Y
	widthCorrection, heightCorrection := 0, 0
	if width%2 != 0 {
		widthCorrection = 1
	}
	if height%2 != 0 {
		heightCorrection = 1
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
			x, y := utilities.GetCoordinates(i/4, width)
			var j int
			skip := true
			if direction == constants.FLIP_DIRECTION_HORIZONTAL && x < width/2+widthCorrection {
				j = utilities.GetPixel(width-x-1, y, width)
				skip = false
			}
			if direction == constants.FLIP_DIRECTION_VERTICAL && y < height/2+heightCorrection {
				j = utilities.GetPixel(x, height-y-1, width)
				skip = false
			}
			if !skip {
				r, g, b := img.Pix[i], img.Pix[i+1], img.Pix[i+2]
				img.Pix[i], img.Pix[i+1], img.Pix[i+2] = img.Pix[j], img.Pix[j+1], img.Pix[j+2]
				img.Pix[j], img.Pix[j+1], img.Pix[j+2] = r, g, b
			}
		}
	}
	for t := 0; t < threads; t += 1 {
		wg.Add(1)
		go processing(t)
	}
	wg.Wait()
	return utilities.EncodeResult(img, format)
}
