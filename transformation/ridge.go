package transformation

import (
	"github.com/salamnocap/go-img-kernel/kernel"
	"github.com/salamnocap/go-img-kernel/utils"
)

func RidgeKernel() *kernel.Kernel {
	k := kernel.NewKernel(
		3,
		[][]float64{
			{-1, -1, -1},
			{-1, 8, -1},
			{-1, -1, -1},
		},
	)
	return k
}

func RidgeDetectionGray(grayImage [][]float64, padding int, stride int) [][]float64 {
	ridge := RidgeKernel()
	return kernel.Convolve2D(grayImage, ridge, padding, stride)
}

func RidgeDetectionRGB(rgbImage [][][]float64, padding int, stride int) [][]float64 {
	grayImage := utils.RgbToGray(rgbImage)
	return RidgeDetectionGray(grayImage, padding, stride)
}
