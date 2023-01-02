package utilities

import (
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"io"
)

func PrepareSource(file io.Reader) ([][]color.Color, string, error) {
	content, format, decodingError := image.Decode(file)
	if decodingError != nil {
		return nil, "", decodingError
	}

	rect := content.Bounds()
	height, width := rect.Dy(), rect.Dx()
	rgba := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(rgba, rgba.Bounds(), content, rect.Min, draw.Src)

	grid := make([][]color.Color, width)
	for x := 0; x < width; x += 1 {
		col := make([]color.Color, height)
		for y := 0; y < height; y += 1 {
			col[y] = rgba.At(x, y)
		}
		grid[x] = col
	}

	return grid, format, nil
}
