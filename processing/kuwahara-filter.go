package processing

import (
	"image/color"

	"github.com/julyskies/brille/utilities"
)

func getAperture(axisValue, axisMax, apertureMin, apertureMax int) (int, int) {
	start, end := 0, axisMax
	if axisValue+apertureMin > 0 {
		start = axisValue + apertureMin
	}
	if axisValue+apertureMax < axisMax {
		end = axisValue + apertureMax
	}
	return start, end
}

func KuwaharaFilter(source [][]color.Color, aperture uint) [][]color.Color {
	width, height := len(source), len(source[0])
	destination := utilities.CreateGrid(width, height)
	apertureHalf := int(aperture / 2)
	apertureMinX := [4]int{-apertureHalf, 0, -apertureHalf, 0}
	apertureMaxX := [4]int{0, apertureHalf, 0, apertureHalf}
	apertureMinY := [4]int{-apertureHalf, -apertureHalf, 0, 0}
	apertureMaxY := [4]int{0, 0, apertureHalf, apertureHalf}
	for x := 0; x < width; x += 1 {
		for y := 0; y < height; y += 1 {
			pixelCount := [4]int{0, 0, 0, 0}
			rValues := [4]int{0, 0, 0, 0}
			gValues := [4]int{0, 0, 0, 0}
			bValues := [4]int{0, 0, 0, 0}
			maxRValue := [4]int{0, 0, 0, 0}
			maxGValue := [4]int{0, 0, 0, 0}
			maxBValue := [4]int{0, 0, 0, 0}
			minRValue := [4]int{255, 255, 255, 255}
			minGValue := [4]int{255, 255, 255, 255}
			minBValue := [4]int{255, 255, 255, 255}
			for i := 0; i < 4; i += 1 {
				x2start, x2end := getAperture(x, width, apertureMinX[i], apertureMaxX[i])
				y2start, y2end := getAperture(y, height, apertureMinY[i], apertureMaxY[i])
				for x2 := x2start; x2 < x2end; x2 += 1 {
					for y2 := y2start; y2 < y2end; y2 += 1 {
						r, g, b, _ := utilities.RGBA(source[x2][y2])
						rValues[i] += int(r)
						gValues[i] += int(g)
						bValues[i] += int(b)
						if int(r) > maxRValue[i] {
							maxRValue[i] = int(r)
						} else if int(r) < minRValue[i] {
							minRValue[i] = int(r)
						}
						if int(g) > maxGValue[i] {
							maxGValue[i] = int(g)
						} else if int(g) < minGValue[i] {
							minGValue[i] = int(g)
						}
						if int(b) > maxBValue[i] {
							maxBValue[i] = int(b)
						} else if int(b) < minBValue[i] {
							minBValue[i] = int(b)
						}
						pixelCount[i] += 1
					}
				}
			}
			j := 0
			MinDifference := 10000
			for i := 0; i < 4; i += 1 {
				cdR := maxRValue[i] - minRValue[i]
				cdG := maxGValue[i] - minGValue[i]
				cdB := maxBValue[i] - minBValue[i]
				CurrentDifference := cdR + cdG + cdB
				if CurrentDifference < MinDifference && pixelCount[i] > 0 {
					j = i
					MinDifference = CurrentDifference
				}
			}
			cR := uint8(rValues[j] / pixelCount[j])
			cG := uint8(gValues[j] / pixelCount[j])
			cB := uint8(bValues[j] / pixelCount[j])
			destination[x][y] = color.RGBA{cR, cG, cB, 255}
		}
	}
	return destination
}
