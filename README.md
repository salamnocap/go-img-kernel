# GO-IMG-KERNEL - [![Godoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/salamnocap/go-img-kernel)
**GO-IMG-KERNEL** is a Go library for image processing, 
leveraging convolutional operations to perform various transformations such as blurring, 
edge detection, binarization, and more. 

This library aims to provide simple-to-use functions for handling images,
whether in RGB or grayscale formats.

## Features

- Image loading and saving in various formats (JPEG, PNG).
- Support for RGB and grayscale images.
- Common image transformations:
  - Gaussian blur
  - Edge detection
  - Binarization
  - Image rotation and reflection
  - Horizontal transformation
  - Sharpening

## Installation
    
```bash
go get github.com/salamnocap/go-img-kernel
```

## Run Tests

```bash
go test -v ./test
```

## Usage
    
```go
import (
  "github.com/salamnocap/go-img-kernel/utils"
  "github.com/salamnocap/go-img-kernel/transformation"
)

rgbImage, err := utils.LoadRGBImage(
    "../examples/input/castle_rgb.jpg",
) // Load RGB image
if err != nil {
    panic(err)
}
newImage := transformation.BoxBlurRGB(rgbImage, 0, 3) // Apply box blur (0 - padding, 3 -stride)
err = utils.SaveImage(
    newImage,
    "../examples/output/blurred.jpg",
    false,
) // Save image
if err != nil {
    panic(err)
}
```

## Custom Kernel

```go
import (
	"github.com/salamnocap/go-img-kernel/kernel"
    "github.com/salamnocap/go-img-kernel/utils"
)

// Custom kernel
customKernel := kernel.NewKernel(
	3, // Kernel size
    [][]float64{
        {0, -1, 0},
        {-1, 4, -1},
        {0, -1, 0},
    },
)

inputImage, err := utils.LoadRGBImage(
    "../examples/input/castle_rgb.jpg",
) // Load RGB image
if err != nil {
    panic(err)
}

outputImage := kernel.Convolve3D(inputImage, customKernel, 0, 1) // Apply custom kernel (0 - padding, 1 - stride)
err = utils.SaveImage(
    newImage,
    "../examples/output/new_image.jpg",
    false, // if false, save as RGB image
) // Save image
if err != nil {
	panic(err)
}
```

## Examples

| Transformation            |                                                                              Kernel Matrix                                                                               | Input Image                                             |                         Output Image                          |
|:--------------------------|:------------------------------------------------------------------------------------------------------------------------------------------------------------------------:|:-------------------------------------------------------:|:-------------------------------------------------------------:|
| Box Blur                  |     $`\begin{bmatrix} \frac{1}{9} & \frac{1}{9} & \frac{1}{9} \\ \frac{1}{9} & \frac{1}{9} & \frac{1}{9} \\ \frac{1}{9} & \frac{1}{9} & \frac{1}{9} \end{bmatrix}`$      | ![castle_rgb.jpg](examples/input/castle_rgb.jpg)         |          ![blurred.jpg](examples/output/blurred.jpg)          |
| Gaussian Blur             | $`\begin{bmatrix} \frac{1}{16} & \frac{2}{16} & \frac{1}{16} \\ \frac{2}{16} & \frac{4}{16} & \frac{2}{16} \\ \frac{1}{16} & \frac{2}{16} & \frac{1}{16} \end{bmatrix}`$ | ![castle_rgb.jpg](examples/input/castle_rgb.jpg) | ![gaussian_blurred.jpg](examples/output/gaussian_blurred.jpg) |
| Edge Detection            |                                                $`\begin{bmatrix} 0 & -1 & 0 \\ -1 & 4 & -1 \\ 0 & -1 & 0 \end{bmatrix}`$                                                 | ![castle_rgb.jpg](examples/input/castle_rgb.jpg)         |        ![edge_detected.jpg](examples/output/edge.jpg)         |
| Diagonal Edge Detection   |                                                 $`\begin{bmatrix} 1 & 0 & -1 \\ 0 & 0 & 0 \\ -1 & 0 & 1 \end{bmatrix}`$                                                  | ![castle_rgb.jpg](examples/input/castle_rgb.jpg)         |        ![edge_detected.jpg](examples/output/edge.jpg)         |
| Binarization              |                                                $`\begin{bmatrix} 0 & -1 & 0 \\ -1 & 4 & -1 \\ 0 & -1 & 0 \end{bmatrix}`$                                                 | ![castle_rgb.jpg](examples/input/castle_rgb.jpg)         |         ![binarized.jpg](examples/output/binary.jpg)          |
| Image Rotation (45°)      |                                                                                  $$-$$                                                                                   | ![castle_rgb.jpg](examples/input/castle_rgb.jpg) |         ![rotated.jpg](examples/output/rotation.jpg)          |
| Image Reflection          |                                                  $`\begin{bmatrix} -1 & 0 & 0 \\ 0 & 1 & 0 \\ 0 & 0 & 1 \end{bmatrix}`$                                                  | ![castle_rgb.jpg](examples/input/castle_rgb.jpg)         |       ![reflected.jpg](examples/output/reflection.jpg)        |
| Horizontal Transformation |                                            $`\begin{bmatrix} -10 & 10 & -10 \\ 10 & 10 & 10 \\ -10 & 10 & -10 \end{bmatrix}`$                                            | ![castle_rgb.jpg](examples/input/castle_rgb.jpg)         | ![horizontal_transformed.jpg](examples/output/horizontal.jpg) |
| Ridge Detection           |                                              $`\begin{bmatrix} -1 & -1 & -1 \\ -1 & 8 & -1 \\ -1 & -1 & -1 \end{bmatrix}`$                                               | ![castle_rgb.jpg](examples/input/castle_rgb.jpg)         |            ![ridge.jpg](examples/output/ridge.jpg)            |
| Sharpening                |                                                $`\begin{bmatrix} 0 & -1 & 0 \\ -1 & 5 & -1 \\ 0 & -1 & 0 \end{bmatrix}`$                                                 | ![castle_rgb.jpg](examples/input/castle_rgb.jpg)         |        ![sharpened.jpg](examples/output/sharpened.jpg)        |
