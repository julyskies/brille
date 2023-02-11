package filters

import (
	"image/color"
	"io"

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
	destination := make([][]color.Color, gridWidth)
	for i := range destination {
		destination[i] = make([]color.Color, gridHeight)
	}
	for i := 0; i < len(img.Pix); i += 4 {
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
	return utilities.EncodeGridResult(destination, format)
}
