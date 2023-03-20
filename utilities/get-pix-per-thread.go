package utilities

import "math"

func GetPixPerThread(pixLen, threads int) int {
	pixPerThreadRaw := float64(pixLen) / float64(threads)
	module := math.Mod(pixPerThreadRaw, 4.0)
	if module == 0 {
		return int(pixPerThreadRaw)
	}
	return int(pixPerThreadRaw + (float64(threads) - math.Mod(pixPerThreadRaw, 4.0)))
}
