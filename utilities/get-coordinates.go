package utilities

import "math"

func GetCoordinates(pixel, width int) (int, int) {
	return pixel % width, int(math.Floor(float64(pixel) / float64(width)))
}
