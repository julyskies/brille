package filters

import (
	"io"

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
	for i := 0; i < len(img.Pix); i += 4 {
		x, y := utilities.GetCoordinates(i/4, width)
		var j int
		if direction == constants.FLIP_DIRECTION_HORIZONTAL && x < width/2+widthCorrection {
			j = utilities.GetPixel(width-x-1, y, width)
		}
		if direction == constants.FLIP_DIRECTION_VERTICAL && y < height/2+heightCorrection {
			j = utilities.GetPixel(x, height-y-1, width)
		}
		r, g, b := img.Pix[i], img.Pix[i+1], img.Pix[i+2]
		img.Pix[i], img.Pix[i+1], img.Pix[i+2] = img.Pix[j], img.Pix[j+1], img.Pix[j+2]
		img.Pix[j], img.Pix[j+1], img.Pix[j+2] = r, g, b
	}
	return utilities.EncodeResult(img, format)
}
