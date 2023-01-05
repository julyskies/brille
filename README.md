## Brille Image Processing

Brille is a small image processing module for Golang

It works with **JPEG** and **PNG** images

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

*// TODO: list of available filters*

### License

[MIT](./LICENSE.md)
