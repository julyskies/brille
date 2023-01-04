package processing

import (
	"image/color"
	"math"

	"github.com/julyskies/brille/utilities"
)

func getPoints(current, amount, total int) (int, int) {
	start, end := 0, total
	if current >= amount {
		start = current - amount
	}
	if current < total-amount {
		end = current + amount
	}
	return start, end
}

func BoxBlur(source [][]color.Color, amount uint) [][]color.Color {
	width, height := len(source), len(source[0])
	destination := utilities.CreateGrid(width, height)
	min := math.Min(float64(width), float64(height))
	if amount > (uint(min) / 2) {
		amount = uint(min / 2)
	}
	amountInt := int(amount)
	var denominator uint
	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			var tR, tG, tB uint
			_, _, _, A := source[x][y].RGBA()
			denominator = 0
			iStart, iEnd := getPoints(x, amountInt, width)
			jStart, jEnd := getPoints(y, amountInt, height)
			for i := iStart; i < iEnd; i += 1 {
				for j := jStart; j < jEnd; j += 1 {
					denominator += 1
					R, G, B, _ := source[i][j].RGBA()
					tR += uint(uint8(R))
					tG += uint(uint8(G))
					tB += uint(uint8(B))
				}
			}
			bR := tR / denominator
			bG := tG / denominator
			bB := tB / denominator
			destination[x][y] = color.RGBA{uint8(bR), uint8(bG), uint8(bB), uint8(A)}
		}
	}
	return destination
}
