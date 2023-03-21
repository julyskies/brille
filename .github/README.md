## Brille Image Processing

Brille is an image processing module for Golang

Works with **JPEG** and **PNG** images

It does not have any dependencies, only standard Golang libraries are being used

Demo: https://images.dyum.in

### Installation

Minimal required Golang version: **v1.18**

```shell script
go get github.com/julyskies/brille/v2@latest
```

### Versioning

Brille `v2.0.X` have a number of breaking changes compared to `v1.0.X`:
- some of the filtering functions were removed
- some of the filtering functions were renamed
- some of the constants were renamed

Check `Available filters` section for more information regarding these changes.

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

	sobel, format, processingError := brille.Sobel(file)
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
		brille.GRAYSCALE_TYPE_LUMINANCE,
	)
	if processingError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	context.Set("Content-Type", "image/"+format)
	return context.SendStream(result)
}
```

Full Fiber example is available at https://github.com/peterdee/filtering-backend

### Performance

Brille uses `sync.WaitGroup` for performance optimization starting with version **v2.0.4**. Image processing can be done in several threads (by default Brille uses `runtime.NumCPU()` value). 

### Available filters

- **Binary**: converts an image to 1-bit black and white. Requires a threshold value (`uint8`, 0 to 255):

  ```golang
  binary, format, processingError := brille.Binary(file, 155)
  ```

- **Box blur**: blurs the image. Requires blur radius to be provided (`uint`, any value):

  ```golang
  blurred, format, processingError := brille.BoxBlur(file, 7)
  ```

- **Brightness**: adjusts image bightness. Requires brightness amount to be provided (`int`, -255 to 255):

  ```golang
  brightened, format, processingError := brille.Brightness(file, 75)
  ```

- **Color inversion**: inverts image colors:

  ```golang
  inverted, format, processingError := brille.ColorInversion(file)
  ```

- **Contrast**: adjusts image contrast. Requires contrast amount to be provided (`int`, -255 to 255):

  ```golang
  contrast, format, processingError := brille.Contrast(file, -45)
  ```

- **Eight colors**: this filter leaves only eight colors present on the image (red, green, blue, yellow, cyan, magenta, white, black):

  ```golang
  eightColors, format, processingError := brille.EightColors(file)
  ```

- **Emboss**: a static edge detection filter that uses a 3x3 kernel. It can be used to outline edges on an image:

  ```golang
  embossed, format, processingError := brille.Emboss(file)
  ```

- **Flip**: flips image horizontally or vertically. This filter requires a second argument - flip direction. Flip directions are available as `brille` module constants (FLIP_DIRECTION_HORIZONTAL is a horizontal mirroring and FLIP_DIRECTION_VERTICAL is vertical mirroring):

  ```golang
  flipped, format, processingError := brille.Flip(
    file,
    brille.FLIP_DIRECTION_HORIZONTAL,
  )
  ```

- **Gamma correction**: corrects image gamma. Requires correction amount to be provided (`float64`, 0 to 3.99). By default image gamma equals to 1, so providing a value less than that makes colors more intense, and larger values decrease color intensity:

  ```golang
  corrected, format, processingError := brille.GammaCorrection(file, 2.05)
  ```

- **Gaussian blur**: blurs an image using dynamically generated Gaussian kernel. Requires sigma value to be provided (`float64`, 0 to 99). Sigma value represents blur intensity and directly correlates with generated kernel size. This Gaussian blur implementation performs 2 convolution cycles (horizontal & vertical) for the best blur quality:

  ```golang
  blurred, format, processingError := brille.GaussianBlur(file, 7.5)
  ```

- **Grayscale**: turns colors into shades of gray. This filter requires a second argument - grayscale type. Grayscale types are available as `brille` module constants (GRAYSCALE_TYPE_AVERAGE and GRAYSCALE_TYPE_LUMINANCE):

  ```golang
  grayscale, format, processingError := brille.Grayscale(
    file,
    brille.GRAYSCALE_TYPE_LUMINANCE,
  )
  ```

- **Hue rotation**: rotates image hue. Requires an angle to be provided (`int`, any value):

  ```golang
  rotated, format, processingError := brille.HueRotate(file, 278)
  ```

- **Kuwahara**: an edge detection filter with dynamic radius. Requires radius to be provided (`uint`, any value). This filter is pretty slow, and will probably be optimized in the future:

  ```golang
  kuwahara, format, processingError := brille.Kuwahara(file, 5)
  ```

- **Laplacian**: a static edge detection filter that uses a 3x3 kernel. It can be used to outline edges on an image:

  ```golang
  laplacian, format, processingError := brille.Laplacian(file)
  ```

- **Rotate image (fixed angles)**: rotates an image clockwise (90, 180 or 270 degrees). This filter requires a second argument - rotation angle. Rotation angles are available as `brille` module constants (ROTATE_FIXED_90, ROTATE_FIXED_180 and ROTATE_FIXED_270):

  ```golang
  rotated270deg, format, processingError := brille.RotateFixed(
    file,
    brille.ROTATE_FIXED_270,
  )
  ```

- **Sepia**: sepia color filter.

  ```golang
  sepia, format, processingError := brille.Sepia(file)
  ```

- **Sharpen**: sharpens provided image. Requires an amount to be provided (`uint`, 0 to 100):

  ```golang
  sharpen, format, processingError := brille.Sharpen(file, 50)
  ```

- **Sobel**: a static edge detection filter that uses a 3x3 kernel. It can be used to outline edges on an image:

  ```golang
  sobel, format, processingError := brille.Sobel(file)
  ```

- **Solarize**: solarization affects image colors, partially inversing the colors. Requires a threshold to be provided (`uint8`, 0 to 255):

  ```golang
  solarized, format, processingError := brille.Solarize(file, 99)
  ```

### Environment variables

- `BRILLE_JPEG_QUALITY` (`int`) - controls output quality for JPEG images, should be a number from 0 (low quality) to 100 (highest quality). Highest quality is used by default.

- `BRILLE_THREADS` (`int`) - controls the number of threads used when performing image processing. By default `runtime.NumCPU()` value is used, and provided number should be less or equal to that value.

### Contributing

Please check [contributing rules](CONTRIBUTING.md).

### License

[MIT](./LICENSE.md)
