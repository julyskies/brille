package utilities

func GetAperture(axisValue, axisMax, apertureMin, apertureMax int) (int, int) {
	start, end := 0, axisMax
	if axisValue+apertureMin > 0 {
		start = axisValue + apertureMin
	}
	if axisValue+apertureMax < axisMax {
		end = axisValue + apertureMax
	}
	return start, end
}
