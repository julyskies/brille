package utilities

import (
	"image/color"
	"math"

	"brille/constants"
)

func Gray(pixel color.Color, calculationType string) (gray uint8, alpha uint8) {
	R, G, B, A := pixel.RGBA()
	alpha = uint8(A)
	if calculationType == constants.GRAYSCALE_LUMINOCITY {
		gray = uint8(
			math.Round(
				(float64(uint8(R))*0.21 + float64(uint8(G))*0.72 + float64(uint8(B))*0.07),
			),
		)
		return
	}
	gray = uint8(
		math.Round(
			(float64(uint8(R)) + float64(uint8(G)) + float64(uint8(B))) / 3.0,
		),
	)
	return
}
