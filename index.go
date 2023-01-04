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

func Binary(file io.Reader, threshold uint) (io.Reader, string, error) {
	if file == nil {
		return nil, "", errors.New(constants.ERROR_NO_FILE_PROVIDED)
	}
	if threshold > 255 {
		return nil, "", errors.New(constants.ERROR_INVALID_BINARY_THRESHOLD)
	}
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
