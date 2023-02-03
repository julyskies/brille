package processing

import (
	"image/color"

	"github.com/julyskies/brille/constants"
	"github.com/julyskies/brille/utilities"
)

var laplacianKernel = [3][3]int{
	{-1, -1, -1},
	{-1, 8, -1},
	{-1, -1, -1},
}

func LaplasianFilter(source [][]color.Color) [][]color.Color {
	width, height := len(source), len(source[0])
	destination := utilities.CreateGrid(width, height)
	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			averageSum := 0
			for i := 0; i < 3; i += 1 {
				for j := 0; j < 3; j += 1 {
					k := utilities.GradientPoint(x, i, width)
					l := utilities.GradientPoint(y, j, height)
					grayColor, _ := utilities.Gray(
						source[x+k][y+l],
						constants.GRAYSCALE_AVERAGE,
					)
					averageSum += int(grayColor) * laplacianKernel[i][j]
				}
			}
			channel := 255 - uint8(utilities.MaxMin(averageSum, 255, 0))
			destination[x][y] = color.RGBA{channel, channel, channel, 255}
		}
	}
	return destination
}
