package transformation

import (
	"github.com/salamnocap/go-img-kernel/kernel"
)

func ReflectionKernel() *kernel.Kernel {
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
	identity := ReflectionKernel()
	return kernel.Convolve3D(rgbImage, identity, padding, stride)
}

func ReflectionTransformationGray(grayImage [][]float64, padding int, stride int) [][]float64 {
	identity := ReflectionKernel()
	return kernel.Convolve2D(grayImage, identity, padding, stride)
}
