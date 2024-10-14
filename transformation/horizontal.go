package transformation

import (
	"github.com/salamnocap/go-img-kernel/kernel"
)

func HorizontalKernel() *kernel.Kernel {
	k := kernel.NewKernel(
		3,
		[][]float64{
			{-10, 10, -10},
			{10, 10, 10},
			{-10, 10, -10},
		},
	)
	return k
}

func HorizontalKernel5() *kernel.Kernel {
	k := kernel.NewKernel(
		5,
		[][]float64{
			{-10, 10, -10},
			{10, 10, 10},
			{-10, 10, -10},
		},
	)
	return k
}

func HorizontalTransformationRGB(rgbImage [][][]float64, padding int, stride int) [][][]float64 {
	horizontal := HorizontalKernel()
	return kernel.Convolve3D(rgbImage, horizontal, padding, stride)
}

func HorizontalTransformationGray(grayImage [][]float64, padding int, stride int) [][]float64 {
	horizontal := HorizontalKernel()
	return kernel.Convolve2D(grayImage, horizontal, padding, stride)
}
