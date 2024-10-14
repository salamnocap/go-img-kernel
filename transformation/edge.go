package transformation

import (
	"Go-Image-Kernels/kernel"
	"Go-Image-Kernels/utils"
)

func edgeKernel() *kernel.Kernel {
	k := kernel.NewKernel(
		3,
		[][]float64{
			{0, -1, 0},
			{-1, 4, -1},
			{0, -1, 0},
		},
	)
	return k
}

func diagonalEdgeKernel() *kernel.Kernel {
	k := kernel.NewKernel(
		3,
		[][]float64{
			{1, 0, -1},
			{0, 0, 0},
			{-1, 0, 1},
		},
	)
	return k
}

func EdgeDetectionRGB(rgbImage [][][]float64, padding int, stride int) [][]float64 {
	edge := edgeKernel()
	grayImage := utils.RgbToGray(rgbImage)
	return kernel.Convolve2D(grayImage, edge, padding, stride)
}

func DiagonalEdgeDetectionRGB(rgbImage [][][]float64, padding int, stride int) [][]float64 {
	edge := diagonalEdgeKernel()
	grayImage := utils.RgbToGray(rgbImage)
	return kernel.Convolve2D(grayImage, edge, padding, stride)
}

func EdgeDetectionGray(grayImage [][]float64, padding int, stride int) [][]float64 {
	edge := edgeKernel()
	return kernel.Convolve2D(grayImage, edge, padding, stride)
}

func DiagonalEdgeDetectionGray(grayImage [][]float64, padding int, stride int) [][]float64 {
	edge := diagonalEdgeKernel()
	return kernel.Convolve2D(grayImage, edge, padding, stride)
}
