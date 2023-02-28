package brille

import (
	"errors"
	"io"

	"github.com/julyskies/brille/v2/constants"
	"github.com/julyskies/brille/v2/filters"
	"github.com/julyskies/brille/v2/utilities"
)

const FLIP_DIRECTION_HORIZONTAL string = constants.FLIP_DIRECTION_HORIZONTAL

const FLIP_DIRECTION_VERTICAL string = constants.FLIP_DIRECTION_VERTICAL

const GRAYSCALE_TYPE_AVERAGE string = constants.GRAYSCALE_TYPE_AVERAGE

const GRAYSCALE_TYPE_LUMINANCE string = constants.GRAYSCALE_TYPE_LUMINANCE

const ROTATE_FIXED_90 uint = constants.ROTATE_FIXED_90

const ROTATE_FIXED_180 uint = constants.ROTATE_FIXED_180

const ROTATE_FIXED_270 uint = constants.ROTATE_FIXED_270

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

// sigma: 0 to 99
func GaussianBlur(file io.Reader, sigma float64) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	return filters.GaussianBlur(file, sigma)
}

// grayscale type: average or luminance
func Grayscale(file io.Reader, grayscaleType string) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	return filters.Grayscale(file, grayscaleType)
}

// angle: any int
func HueRotate(file io.Reader, angle int) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	return filters.HueRotate(file, angle)
}

// radius: 0 to 40
func Kuwahara(file io.Reader, radius uint) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	radius = utilities.Clamp(radius, 40, 0)
	return filters.Kuwahara(file, radius)
}

func Laplacian(file io.Reader) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	return filters.Laplacian(file)
}

// angle: 90, 180 or 270 (use provided constants)
func RotateFixed(file io.Reader, angle uint) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	return filters.RotateFixed(file, angle)
}

func Sepia(file io.Reader) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	return filters.Sepia(file)
}

// amount: 0 to 100
func Sharpen(file io.Reader, amount uint) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	return filters.Sharpen(file, amount)
}

func Sobel(file io.Reader) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	return filters.Sobel(file)
}

// threshold: any uint8
func Solarize(file io.Reader, threshold uint8) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	return filters.Solarize(file, threshold)
}
