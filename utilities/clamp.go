package utilities

func Clamp[T float64 | int | uint](value, max, min T) T {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func ClampMax[T float64 | int | uint](value, max T) T {
	if value > max {
		return max
	}
	return value
}
