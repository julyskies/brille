package filters

import (
	"io"

	"github.com/julyskies/brille/v2/utilities"
)

type Color struct {
	R, G, B int
}

var COLORS = [8]Color{
	{255, 0, 0},
	{0, 255, 0},
	{0, 0, 255},
	{255, 255, 0},
	{255, 0, 255},
	{0, 255, 255},
	{255, 255, 255},
	{0, 0, 0},
}

func EightColors(file io.Reader) (io.Reader, string, error) {
	img, format, convertationError := utilities.DecodeSource(file)
	if convertationError != nil {
		return nil, "", convertationError
	}
	for i := 0; i < len(img.Pix); i += 4 {
		minDelta := 195076
		var selectedColor Color
		for j := range COLORS {
			indexColor := COLORS[j]
			rDifference := int(img.Pix[i]) - indexColor.R
			gDifference := int(img.Pix[i+1]) - indexColor.G
			bDifference := int(img.Pix[i+2]) - indexColor.B
			delta := rDifference*rDifference + gDifference*gDifference + bDifference*bDifference
			if delta < minDelta {
				minDelta = delta
				selectedColor = indexColor
			}
		}
		img.Pix[i] = uint8(selectedColor.R)
		img.Pix[i+1] = uint8(selectedColor.G)
		img.Pix[i+2] = uint8(selectedColor.B)
	}
	return utilities.EncodeResult(img, format)
}
