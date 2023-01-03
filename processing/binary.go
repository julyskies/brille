package processing

import (
	"image/color"

	"github.com/julyskies/brille/constants"
	"github.com/julyskies/brille/utilities"
)

func Binary(source [][]color.Color, threshold uint) [][]color.Color {
	width, height := len(source), len(source[0])
	destination := utilities.CreateGrid(width, height)
	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			grayColor, alpha := utilities.Gray(source[x][y], constants.GRAYSCALE_AVERAGE)
			value := uint8(255)
			if uint(grayColor) < threshold {
				value = 0
			}
			destination[x][y] = color.RGBA{value, value, value, alpha}
		}
	}
	return destination
}
