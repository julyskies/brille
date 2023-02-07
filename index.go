package brille

import (
	"errors"
	"io"

	"github.com/julyskies/brille/constants"
	"github.com/julyskies/brille/filters"
	"github.com/julyskies/brille/processing"
	"github.com/julyskies/brille/utilities"
)

const FLIP_DIRECTION_HORIZONTAL string = constants.FLIP_DIRECTION_HORIZONTAL

const FLIP_DIRECTION_VERTICAL string = constants.FLIP_DIRECTION_VERTICAL

const GRAYSCALE_AVERAGE string = constants.GRAYSCALE_AVERAGE

const GRAYSCALE_LUMINOCITY string = constants.GRAYSCALE_LUMINOCITY

/* Optimized filters */

// threshold: 0 to 255
func Binary(file io.Reader, threshold uint8) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	return filters.Binary(file, threshold)
}

// radius: any uint
func BoxBlur(file io.Reader, radius uint) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	return filters.BoxBlur(file, radius)
}

// amount: -255 to 255
func Brightness(file io.Reader, amount int) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	return filters.Brightness(file, amount)
}

func ColorInversion(file io.Reader) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	return filters.ColorInversion(file)
}

// amount: -255 to 255
func Contrast(file io.Reader, amount int) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	return filters.Contrast(file, amount)
}

func EightColors(file io.Reader) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	return filters.EightColors(file)
}

func Emboss(file io.Reader) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	return filters.Emboss(file)
}

// direction: horizontal or vertical
func Flip(file io.Reader, direction string) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	return filters.Flip(file, direction)
}

// amount: 0 to 3.99
func GammaCorrection(file io.Reader, amount float64) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	return filters.GammaCorrection(file, amount)
}

/* Non-optimized filters */

// type: average or luminocity
func Grayscale(file io.Reader, grayscaleType string) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	if grayscaleType != GRAYSCALE_AVERAGE &&
		grayscaleType != GRAYSCALE_LUMINOCITY {
		return nil, "", errors.New(constants.ERROR_INVALID_GRAYSCALE_TYPE)
	}
	source, format, preparationError := utilities.PrepareSource(file)
	if preparationError != nil {
		return nil, "", preparationError
	}
	grayscale := processing.Grayscale(source, grayscaleType)
	encoded, encodingError := utilities.PrepareResult(grayscale, format)
	if encodingError != nil {
		return nil, "", encodingError
	}
	return encoded, format, nil
}

// angle: any int value
func HueRotate(file io.Reader, angle int) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	source, format, preparationError := utilities.PrepareSource(file)
	if preparationError != nil {
		return nil, "", preparationError
	}
	rotated := processing.HueRotate(source, angle)
	encoded, encodingError := utilities.PrepareResult(rotated, format)
	if encodingError != nil {
		return nil, "", encodingError
	}
	return encoded, format, nil
}

// aperture: 0 to 40
func KuwaharaFilter(file io.Reader, aperture uint) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	aperture = utilities.MaxMin(aperture, 40, 0)
	source, format, preparationError := utilities.PrepareSource(file)
	if preparationError != nil {
		return nil, "", preparationError
	}
	kuwahara := processing.KuwaharaFilter(source, aperture)
	encoded, encodingError := utilities.PrepareResult(kuwahara, format)
	if encodingError != nil {
		return nil, "", encodingError
	}
	return encoded, format, nil
}

func LaplasianFilter(file io.Reader) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	source, format, preparationError := utilities.PrepareSource(file)
	if preparationError != nil {
		return nil, "", preparationError
	}
	laplasian := processing.LaplasianFilter(source)
	encoded, encodingError := utilities.PrepareResult(laplasian, format)
	if encodingError != nil {
		return nil, "", encodingError
	}
	return encoded, format, nil
}

func Rotate90(file io.Reader) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	source, format, preparationError := utilities.PrepareSource(file)
	if preparationError != nil {
		return nil, "", preparationError
	}
	rotated := processing.ImageRotation(source, 90)
	encoded, encodingError := utilities.PrepareResult(rotated, format)
	if encodingError != nil {
		return nil, "", encodingError
	}
	return encoded, format, nil
}

func Rotate180(file io.Reader) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	source, format, preparationError := utilities.PrepareSource(file)
	if preparationError != nil {
		return nil, "", preparationError
	}
	rotated := processing.ImageRotation(source, 180)
	encoded, encodingError := utilities.PrepareResult(rotated, format)
	if encodingError != nil {
		return nil, "", encodingError
	}
	return encoded, format, nil
}

func Rotate270(file io.Reader) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	source, format, preparationError := utilities.PrepareSource(file)
	if preparationError != nil {
		return nil, "", preparationError
	}
	rotated := processing.ImageRotation(source, 270)
	encoded, encodingError := utilities.PrepareResult(rotated, format)
	if encodingError != nil {
		return nil, "", encodingError
	}
	return encoded, format, nil
}

func Sepia(file io.Reader) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	source, format, preparationError := utilities.PrepareSource(file)
	if preparationError != nil {
		return nil, "", preparationError
	}
	sepia := processing.Sepia(source)
	encoded, encodingError := utilities.PrepareResult(sepia, format)
	if encodingError != nil {
		return nil, "", encodingError
	}
	return encoded, format, nil
}

// amount: 0 to 100
func SharpenFilter(file io.Reader, amount uint) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	mix := float64(utilities.MaxMin(amount, 100, 0)) / 100
	source, format, preparationError := utilities.PrepareSource(file)
	if preparationError != nil {
		return nil, "", preparationError
	}
	sharpen := processing.Sharpen(source, mix)
	encoded, encodingError := utilities.PrepareResult(sharpen, format)
	if encodingError != nil {
		return nil, "", encodingError
	}
	return encoded, format, nil
}

func SobelFilter(file io.Reader) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	source, format, preparationError := utilities.PrepareSource(file)
	if preparationError != nil {
		return nil, "", preparationError
	}
	sobel := processing.SobelFilter(source)
	encoded, encodingError := utilities.PrepareResult(sobel, format)
	if encodingError != nil {
		return nil, "", encodingError
	}
	return encoded, format, nil
}

// threshold: 0 to 255
func Solarize(file io.Reader, threshold uint) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	threshold = utilities.MaxMin(threshold, 255, 0)
	source, format, preparationError := utilities.PrepareSource(file)
	if preparationError != nil {
		return nil, "", preparationError
	}
	solarized := processing.Solarize(source, threshold)
	encoded, encodingError := utilities.PrepareResult(solarized, format)
	if encodingError != nil {
		return nil, "", encodingError
	}
	return encoded, format, nil
}
