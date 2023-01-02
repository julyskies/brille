package brille

import (
	"brille/processing"
	"brille/utilities"
	"io"
)

func Grayscale(file io.Reader, grayscaleType string) (io.Reader, string, error) {
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
