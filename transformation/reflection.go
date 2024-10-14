package transformation

import (
	"Go-Image-Kernels/kernel"
)

func reflectionKernel() *kernel.Kernel {
	k := kernel.NewKernel(
		3,
		[][]float64{
			{-1, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
		},
	)
	return k
}

func ReflectionTransformationRGB(rgbImage [][][]float64, padding int, stride int) [][][]float64 {
	identity := reflectionKernel()
	return kernel.Convolve3D(rgbImage, identity, padding, stride)
}

func ReflectionTransformationGray(grayImage [][]float64, padding int, stride int) [][]float64 {
	identity := reflectionKernel()
	return kernel.Convolve2D(grayImage, identity, padding, stride)
}
