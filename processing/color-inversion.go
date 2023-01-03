package processing

import (
	"image/color"

	"github.com/julyskies/brille/utilities"
)

func ColorInversion(source [][]color.Color) [][]color.Color {
	width, height := len(source), len(source[0])
	destination := utilities.CreateGrid(width, height)
	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			R, G, B, A := source[x][y].RGBA()
			destination[x][y] = color.RGBA{
				255 - uint8(R),
				255 - uint8(G),
				255 - uint8(B),
				uint8(A),
			}
		}
	}
	return destination
}
