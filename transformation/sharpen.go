package transformation

import (
	"github.com/salamnocap/go-img-kernel/kernel"
)

func SharpenKernel() *kernel.Kernel {
	k := kernel.NewKernel(
		3,
		[][]float64{
			{0, -1, 0},
			{-1, 5, -1},
			{0, -1, 0},
		},
	)
	return k
}

func SharpenImageRGB(rgbImage [][][]float64, padding int, stride int) [][][]float64 {
	sharpen := SharpenKernel()
	return kernel.Convolve3D(rgbImage, sharpen, padding, stride)
}

func SharpenImageGray(grayImage [][]float64, padding int, stride int) [][]float64 {
	sharpen := SharpenKernel()
	return kernel.Convolve2D(grayImage, sharpen, padding, stride)
}
