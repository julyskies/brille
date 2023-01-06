## Brille Image Processing

Brille is an image processing module for Golang

Works with **JPEG** and **PNG** images

It does not have any dependencies - only standard Golang libraries are being used

### Installation

Minimal required Golang version: **v1.18**

```shell script
go get github.com/julyskies/brille@latest
```

### Usage

Basic example:

```golang
package main

import (
  "log"
  "os"

  "github.com/julyskies/brille"
)

func main() {
  file, _ := os.Open("/path/to/file.jpeg")
  defer file.Close()

  grayscaled, fileFormat, processingError := brille.Grayscale(
    file,
    brille.GRAYSCALE_LUMINOCITY,
  )
  if processingError != nil {
    log.Fatal(processingError)
  }
  // TODO: finish
}
```

Using with Fiber:

```golang
// TODO: Fiber example
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

### License

[MIT](./LICENSE.md)
