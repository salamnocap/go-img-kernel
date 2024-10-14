package transformation

import (
	"Go-Image-Kernels/kernel"
)

func horizontalKernel() *kernel.Kernel {
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

func horizontalKernel5() *kernel.Kernel {
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
	horizontal := horizontalKernel()
	return kernel.Convolve3D(rgbImage, horizontal, padding, stride)
}

func HorizontalTransformationGray(grayImage [][]float64, padding int, stride int) [][]float64 {
	horizontal := horizontalKernel()
	return kernel.Convolve2D(grayImage, horizontal, padding, stride)
}
