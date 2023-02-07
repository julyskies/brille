package utilities

func GetPixel(x, y, width int) int {
	return ((y * width) + x) * 4
}
