package processing

import (
	"image/color"
	"math"

	"github.com/julyskies/brille/constants"
	"github.com/julyskies/brille/utilities"
)

var laplasianKernel = [3][3]int{
	{0, -1, 0},
	{-1, 4, -1},
	{0, -1, 0},
}

func getGradientPoint(axisValue, shift, axisLength int) int {
	if axisValue+shift >= axisLength {
		return axisLength - axisValue - 1
	}
	return shift
}

func LaplasianFilter(source [][]color.Color) [][]color.Color {
	width, height := len(source), len(source[0])
	destination := utilities.CreateGrid(width, height)
	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			gradientX := 0
			for i := 0; i < 3; i += 1 {
				for j := 0; j < 3; j += 1 {
					k := getGradientPoint(x, i, width)
					l := getGradientPoint(y, j, height)
					grayColor, _ := utilities.Gray(
						source[x+k][y+l],
						constants.GRAYSCALE_AVERAGE,
					)
					gradientX += int(grayColor) * laplasianKernel[i][j]
				}
			}
			colorCode := 255 - uint8(int(math.Sqrt(
				float64((gradientX * gradientX)),
			)))
			destination[x][y] = color.RGBA{colorCode, colorCode, colorCode, 255}
		}
	}
	return destination
}
