package main

import (
	"Go-Image-Kernels/transformation"
	"Go-Image-Kernels/utils"
)

func main() {
	rgbImage, err := utils.LoadRGBImage(
		"../examples/input/castle_rgb.jpg",
	)
	if err != nil {
		panic(err)
	}
	newImage := transformation.BoxBlurRGB(rgbImage, 0, 3)
	err = utils.SaveImage(
		newImage,
		"../examples/output/blurred.jpg",
		false,
	)
	if err != nil {
		panic(err)
	}
}
