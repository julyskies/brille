package processing

import (
	"image/color"

	"github.com/julyskies/brille/utilities"
)

func ImageRotation(source [][]color.Color, angle uint) [][]color.Color {
	width, height := len(source), len(source[0])
	var destination [][]color.Color
	if angle == 90 || angle == 270 {
		destination = utilities.CreateGrid(height, width)
	} else {
		destination = utilities.CreateGrid(width, height)
	}
	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			if angle == 90 {
				destination[height-y-1][x] = source[x][y]
			}
			if angle == 180 {
				destination[width-x-1][height-y-1] = source[x][y]
			}
			if angle == 270 {
				destination[y][width-x-1] = source[x][y]
			}
		}
	}
	return destination
}
