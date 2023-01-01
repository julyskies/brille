package brille

import (
	"brille/processing"
	"brille/utilities"
	"io"
)

func Grayscale(file io.Reader, grayscaleType string) (io.Reader, string, error) {
	source, format, preparationError := utilities.PrepareImage(file)
	if preparationError != nil {
		return nil, "", preparationError
	}
	result := processing.Grayscale(source, grayscaleType)
}
