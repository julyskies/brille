package processing

import (
	"image/color"
	"math"

	"github.com/julyskies/brille/constants"
	"github.com/julyskies/brille/utilities"
)

var sobelHorizontal = [3][3]int{
	{-1, 0, 1},
	{-2, 0, 2},
	{-1, 0, 1},
}

var sobelVertical = [3][3]int{
	{1, 2, 1},
	{0, 0, 0},
	{-1, -2, -1},
}

func SobelFilter(source [][]color.Color) [][]color.Color {
	width, height := len(source), len(source[0])
	destination := utilities.CreateGrid(width, height)
	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			gradientX := 0
			gradientY := 0
			for i := 0; i < 3; i += 1 {
				for j := 0; j < 3; j += 1 {
					k := utilities.GradientPoint(x, i, width)
					l := utilities.GradientPoint(y, j, height)
					grayColor, _ := utilities.Gray(
						source[x+k][y+l],
						constants.GRAYSCALE_AVERAGE,
					)
					gradientX += int(grayColor) * sobelHorizontal[i][j]
					gradientY += int(grayColor) * sobelVertical[i][j]
				}
			}
			colorCode := 255 - uint8(utilities.MaxMin(math.Sqrt(
				float64(gradientX*gradientX+gradientY*gradientY),
			), 255, 0))
			destination[x][y] = color.RGBA{colorCode, colorCode, colorCode, 255}
		}
	}
	return destination
}
