package processing

import (
	"image/color"

	"github.com/julyskies/brille/utilities"
)

var sharpenKernel = [3][3]int{
	{-1, -1, -1},
	{-1, 9, -1},
	{-1, -1, -1},
}

func Sharpen(source [][]color.Color, mix float64) [][]color.Color {
	width, height := len(source), len(source[0])
	destination := utilities.CreateGrid(width, height)
	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			sumR := 0
			sumG := 0
			sumB := 0
			for i := 0; i < 3; i += 1 {
				for j := 0; j < 3; j += 1 {
					k := utilities.GradientPoint(x, i, width)
					l := utilities.GradientPoint(y, j, height)
					r, g, b, _ := utilities.RGBA(source[x+k][y+l])
					sumR += int(r) * sharpenKernel[i][j]
					sumG += int(g) * sharpenKernel[i][j]
					sumB += int(b) * sharpenKernel[i][j]
				}
			}
			r, g, b, alpha := utilities.RGBA(source[x][y])
			R := utilities.MaxMin(float64(sumR)*mix+float64(r)*(1-mix), 255, 0)
			G := utilities.MaxMin(float64(sumG)*mix+float64(g)*(1-mix), 255, 0)
			B := utilities.MaxMin(float64(sumB)*mix+float64(b)*(1-mix), 255, 0)
			destination[x][y] = color.RGBA{uint8(R), uint8(G), uint8(B), alpha}
		}
	}
	return destination
}
