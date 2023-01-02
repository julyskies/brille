package processing

import (
	"image/color"

	"brille/constants"
	"brille/utilities"
)

func Grayscale(source [][]color.Color, grayscaleType string) [][]color.Color {
	width, height := len(source), len(source[0])
	destination := utilities.CreateGrid(width, height)
	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			var grayColor, alpha uint8
			if grayscaleType == constants.GRAYSCALE_LUMINOCITY {
				grayColor, alpha = utilities.Gray(source[x][y], constants.GRAYSCALE_LUMINOCITY)
			} else {
				grayColor, alpha = utilities.Gray(source[x][y], constants.GRAYSCALE_AVERAGE)
			}
			destination[x][y] = color.RGBA{grayColor, grayColor, grayColor, alpha}
		}
	}
	return destination
}
