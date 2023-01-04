package utilities

func GradientPoint(axisValue, shift, axisLength int) int {
	if axisValue+shift >= axisLength {
		return axisLength - axisValue - 1
	}
	return shift
}
