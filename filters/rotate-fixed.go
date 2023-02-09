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
	heightCorrection := 0
	if height%2 != 0 {
		heightCorrection = 1
	}
	if angle == constants.ROTATE_FIXED_180 {
		for i := 0; i < len(img.Pix); i += 4 {
			x, y := utilities.GetCoordinates(i/4, width)
			r, g, b := img.Pix[i], img.Pix[i+1], img.Pix[i+2]
			var j int
			if y < height/2+heightCorrection {
				j = utilities.GetPixel(width-x-1, height-y-1, width)
				img.Pix[i], img.Pix[i+1], img.Pix[i+2] = img.Pix[j], img.Pix[j+1], img.Pix[j+2]
				img.Pix[j], img.Pix[j+1], img.Pix[j+2] = r, g, b
			}
		}
		return utilities.EncodeResult(img, format)
	}
	destination := make([][]color.Color, width)
	for i := range destination {
		destination[i] = make([]color.Color, height)
	}
	for i := 0; i < len(img.Pix); i += 4 {
		x, y := utilities.GetCoordinates(i/4, width)
		if angle == constants.ROTATE_FIXED_90 {
			destination[height-y-1][x] = color.RGBA{
				img.Pix[i],
				img.Pix[i+1],
				img.Pix[i+2],
				img.Pix[i+3],
			}
		}
		if angle == constants.ROTATE_FIXED_270 {
			destination[y][width-x-1] = color.RGBA{
				img.Pix[i],
				img.Pix[i+1],
				img.Pix[i+2],
				img.Pix[i+3],
			}
		}
	}
	return utilities.EncodeGridResult(destination, format)
}
