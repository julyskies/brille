## Brille Image Processing

Brille is an image processing module for Golang

Works with **JPEG** and **PNG** images

It does not have any dependencies, only standard Golang libraries are being used

### Installation

Minimal required Golang version: **v1.18**

```shell script
go get github.com/julyskies/brille@latest
```

### Usage

All of the Brille filter functions return 3 values:
- processed file as `io.Reader`
- file format as `string`
- processing error 

Basic example:

```golang
package main

import (
	"io"
	"log"
	"os"

	"github.com/julyskies/brille"
)

func main() {
	file, fileError := os.Open("/path/to/file.jpeg")
	if fileError != nil {
		log.Fatal(fileError)
	}
	defer file.Close()

	sobel, format, processingError := brille.SobelFilter(file)
	if processingError != nil {
		log.Fatal(processingError)
	}

	outputFile, outputError := os.Create("sobel." + format)
	if outputError != nil {
		log.Fatal(outputError)
	}
	defer outputFile.Close()

	_, copyError := io.Copy(outputFile, sobel)
	if copyError != nil {
		log.Fatal(copyError)
	}
	
	log.Println("processed " + format)
}
```

Using with **Fiber**:

```golang
package apis

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/julyskies/brille"
)

func controller(context *fiber.Ctx) error {
	file, fileError := context.FormFile("image")
	if fileError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	partials := strings.Split(file.Filename, ".")
	extension := partials[len(partials)-1]
	if extension == "" {
		return fiber.NewError(fiber.StatusBadRequest)
	}

	fileHandle, readerError := file.Open()
	if readerError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	result, format, processingError := brille.Grayscale(
		fileHandle,
		brille.GRAYSCALE_AVERAGE,
	)
	if processingError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	context.Set("Content-Type", "image/"+format)
	return context.SendStream(result)
}
```

### Available filters

- **Binary**: convert an image to 1 bit black and white. Requires a threshold value (from 0 to 255):

  ```golang
  binary, format, processingError := brille.Binary(file, 155)
  ```

- **Box blur**: blur the image. Requires blur amount to be provided. Blur amount is a `uint` value. Box blur function will automatically reduce the amount value if it is too big, maximum amount is `min(width, height) / 2`:

  ```golang
  blurred, format, processingError := brille.BoxBlur(file, 5)
  ```

- **Brightness**: adjust image bightness. Requires brightness amount to be provided. Brightness amount ranges from -255 (darker) to 255 (brighter):

  ```golang
  brightened, format, processingError := brille.Brightness(file, 107)
  ```

- **Color inversion**: invert image colors.

  ```golang
  inverted, format, processingError := brille.ColorInversion(file)
  ```

- **Contrast**: adjust image contrast. Requires contrast amount to be provided. Contrast amount ranges from -255 (less contrast) to 255 (more contrast):

  ```golang
  contrastless, format, processingError := brille.Contrast(file, -40)
  ```

- **Eight colors**: this filter leaves only eight colors present on the image (red, green, blue, yellow, cyan, magenta, white, black). 

  ```golang
  indexedColors, format, processingError := brille.EightColors(file)
  ```

- **Emboss filter**: a static edge detection filter that uses a 3x3 kernel. It can be used to outline edges on an image:

  ```golang
  embossed, format, processingError := brille.EmbossFilter(file)
  ```

- **Flip horizontal**: flip image horizontally, basically reflect the image in *X* axis.

  ```golang
  flippedX, format, processingError := brille.FlipHorizontal(file)
  ```

- **Flip vertical**: flip image vertically, basically reflect the image in *Y* axis.

  ```golang
  flippedY, format, processingError := brille.FlipVertical(file)
  ```

- **Gamma correction**: image gamma correction. Requires correction amount to be provided. Correction amount ranges from `0` to `3.99` (`float64`). By default image gamma equals to `1`, so providing a value less than `1` makes colors more intense, and values more than `1` decrease color intensity:

  ```golang
  corrected, format, processingError := brille.GammaCorrection(file, 2.05)
  ```

- **Grayscale**: turn colors into shades of gray. Requires grayscale type to be provided. There are 2 grayscale types available: `average` (or `mean`) and `luminocity` (or `weighted`). Both types are available as constants in `brille` module:

  ```golang
  grayAverage, format, processingError := brille.Grayscale(
    file,
    brille.GRAYSCALE_AVERAGE,
  )

  grayLuminocity, format, processingError := brille.Grayscale(
    file,
    brille.GRAYSCALE_LUMINOCITY,
  )
  ```

- **Laplasian filter**: a static edge detection filter that uses a 3x3 kernel. It can be used to outline edges on an image:

  ```golang
  laplasian, format, processingError := brille.LaplasianFilter(file)
  ```

- **Rotate image (fixed angle)**: rotate an image. Available fixed angeles are 90, 180 and 270 degrees (clockwise):

  ```golang
  rotated90, format, processingError := brille.Rotate90(file)

  rotated180, format, processingError := brille.Rotate180(file)
  
  rotated270, format, processingError := brille.Rotate270(file)
  ```

- **Sepia**: sepia color filter.

  ```golang
  sepia, format, processingError := brille.Sepia(file)
  ```

- **Sobel filter**: a static edge detection filter that uses a 3x3 kernel. It can be used to outline edges on an image:

  ```golang
  sobel, format, processingError := brille.Sobel(file)
  ```

- **Solarize**: solarization affects image colors, partially inversing the colors. Requires a threshold to be provided. Threshold ranges from 0 to 255:

  ```golang
  solarized, format, processingError := brille.Solarize(file, 99)
  ```

### License

[MIT](./LICENSE.md)
