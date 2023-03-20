package filters

import (
	"image/color"
	"io"
	"sync"

	"github.com/julyskies/brille/v2/constants"
	"github.com/julyskies/brille/v2/utilities"
)

func RotateFixed(file io.Reader, angle uint) (io.Reader, string, error) {
	img, format, convertationError := utilities.DecodeSource(file)
	if convertationError != nil {
		return nil, "", convertationError
	}
	if angle != constants.ROTATE_FIXED_90 &&
		angle != constants.ROTATE_FIXED_180 &&
		angle != constants.ROTATE_FIXED_270 {
		angle = constants.ROTATE_FIXED_90
	}
	width, height := img.Rect.Max.X, img.Rect.Max.Y
	gridWidth, gridHeight := width, height
	if angle != constants.ROTATE_FIXED_180 {
		gridWidth, gridHeight = height, width
	}
	destination := utilities.CreateGrid(gridWidth, gridHeight)
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
			dx, dy := y, x
			if angle == constants.ROTATE_FIXED_90 {
				dx = height - y - 1
			}
			if angle == constants.ROTATE_FIXED_180 {
				dx, dy = width-x-1, height-y-1
			}
			if angle == constants.ROTATE_FIXED_270 {
				dy = width - x - 1
			}
			destination[dx][dy] = color.RGBA{
				img.Pix[i],
				img.Pix[i+1],
				img.Pix[i+2],
				img.Pix[i+3],
			}
		}
	}
	for t := 0; t < threads; t += 1 {
		wg.Add(1)
		go processing(t)
	}
	wg.Wait()
	return utilities.EncodeGridResult(destination, format)
}
