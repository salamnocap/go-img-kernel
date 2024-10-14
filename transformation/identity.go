package transformation

import (
	"Go-Image-Kernels/kernel"
)

func identityKernel() *kernel.Kernel {
	k := kernel.NewKernel(
		3,
		[][]float64{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		},
	)
	return k
}

func IdentityTransformationRGB(rgbImage [][][]float64, padding int, stride int) [][][]float64 {
	identity := identityKernel()
	return kernel.Convolve3D(rgbImage, identity, padding, stride)
}

func IdentityTransformationGray(grayImage [][]float64, padding int, stride int) [][]float64 {
	identity := identityKernel()
	return kernel.Convolve2D(grayImage, identity, padding, stride)
}
