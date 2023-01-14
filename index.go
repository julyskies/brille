package brille

import (
	"errors"
	"io"

	"github.com/julyskies/brille/constants"
	"github.com/julyskies/brille/processing"
	"github.com/julyskies/brille/utilities"
)

const GRAYSCALE_AVERAGE string = constants.GRAYSCALE_AVERAGE

const GRAYSCALE_LUMINOCITY string = constants.GRAYSCALE_LUMINOCITY

// threshold: 0 to 255
func Binary(file io.Reader, threshold uint) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	threshold = utilities.MaxMin(threshold, 255, 0)
	source, format, preparationError := utilities.PrepareSource(file)
	if preparationError != nil {
		return nil, "", preparationError
	}
	binary := processing.Binary(source, threshold)
	encoded, encodingError := utilities.PrepareResult(binary, format)
	if encodingError != nil {
		return nil, "", encodingError
	}
	return encoded, format, nil
}

// max amount: (min(width, height) / 2)
func BoxBlur(file io.Reader, amount uint) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	source, format, preparationError := utilities.PrepareSource(file)
	if preparationError != nil {
		return nil, "", preparationError
	}
	blurred := processing.BoxBlur(source, amount)
	encoded, encodingError := utilities.PrepareResult(blurred, format)
	if encodingError != nil {
		return nil, "", encodingError
	}
	return encoded, format, nil
}

// amount: from -255 to 255
func Brightness(file io.Reader, amount int) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	amount = utilities.MaxMin(amount, 255, -255)
	source, format, preparationError := utilities.PrepareSource(file)
	if preparationError != nil {
		return nil, "", preparationError
	}
	brightness := processing.Brightness(source, amount)
	encoded, encodingError := utilities.PrepareResult(brightness, format)
	if encodingError != nil {
		return nil, "", encodingError
	}
	return encoded, format, nil
}

func ColorInversion(file io.Reader) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	source, format, preparationError := utilities.PrepareSource(file)
	if preparationError != nil {
		return nil, "", preparationError
	}
	binary := processing.ColorInversion(source)
	encoded, encodingError := utilities.PrepareResult(binary, format)
	if encodingError != nil {
		return nil, "", encodingError
	}
	return encoded, format, nil
}

// amount: from -255 to 255
func Contrast(file io.Reader, amount int) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	amount = utilities.MaxMin(amount, 255, -255)
	source, format, preparationError := utilities.PrepareSource(file)
	if preparationError != nil {
		return nil, "", preparationError
	}
	contrast := processing.Contrast(source, amount)
	encoded, encodingError := utilities.PrepareResult(contrast, format)
	if encodingError != nil {
		return nil, "", encodingError
	}
	return encoded, format, nil
}

func EightColors(file io.Reader) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	source, format, preparationError := utilities.PrepareSource(file)
	if preparationError != nil {
		return nil, "", preparationError
	}
	eightColors := processing.EightColors(source)
	encoded, encodingError := utilities.PrepareResult(eightColors, format)
	if encodingError != nil {
		return nil, "", encodingError
	}
	return encoded, format, nil
}

func EmbossFilter(file io.Reader) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	source, format, preparationError := utilities.PrepareSource(file)
	if preparationError != nil {
		return nil, "", preparationError
	}
	emboss := processing.EmbossFilter(source)
	encoded, encodingError := utilities.PrepareResult(emboss, format)
	if encodingError != nil {
		return nil, "", encodingError
	}
	return encoded, format, nil
}

func FlipHorizontal(file io.Reader) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	source, format, preparationError := utilities.PrepareSource(file)
	if preparationError != nil {
		return nil, "", preparationError
	}
	flipped := processing.FlipHorizontal(source)
	encoded, encodingError := utilities.PrepareResult(flipped, format)
	if encodingError != nil {
		return nil, "", encodingError
	}
	return encoded, format, nil
}

func FlipVertical(file io.Reader) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	source, format, preparationError := utilities.PrepareSource(file)
	if preparationError != nil {
		return nil, "", preparationError
	}
	flipped := processing.FlipVertical(source)
	encoded, encodingError := utilities.PrepareResult(flipped, format)
	if encodingError != nil {
		return nil, "", encodingError
	}
	return encoded, format, nil
}

// amount: from 0 to 3.99
func GammaCorrection(file io.Reader, amount float64) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	amount = utilities.MaxMin(amount, 3.99, 0)
	source, format, preparationError := utilities.PrepareSource(file)
	if preparationError != nil {
		return nil, "", preparationError
	}
	gamma := processing.GammaCorrection(source, amount)
	encoded, encodingError := utilities.PrepareResult(gamma, format)
	if encodingError != nil {
		return nil, "", encodingError
	}
	return encoded, format, nil
}

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
