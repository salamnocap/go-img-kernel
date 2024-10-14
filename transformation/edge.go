package transformation

import (
	"github.com/salamnocap/go-img-kernel/kernel"
	"github.com/salamnocap/go-img-kernel/utils"
)

func EdgeKernel() *kernel.Kernel {
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

func DiagonalEdgeKernel() *kernel.Kernel {
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
	edge := EdgeKernel()
	grayImage := utils.RgbToGray(rgbImage)
	return kernel.Convolve2D(grayImage, edge, padding, stride)
}

func DiagonalEdgeDetectionRGB(rgbImage [][][]float64, padding int, stride int) [][]float64 {
	edge := DiagonalEdgeKernel()
	grayImage := utils.RgbToGray(rgbImage)
	return kernel.Convolve2D(grayImage, edge, padding, stride)
}

func EdgeDetectionGray(grayImage [][]float64, padding int, stride int) [][]float64 {
	edge := EdgeKernel()
	return kernel.Convolve2D(grayImage, edge, padding, stride)
}

func DiagonalEdgeDetectionGray(grayImage [][]float64, padding int, stride int) [][]float64 {
	edge := DiagonalEdgeKernel()
	return kernel.Convolve2D(grayImage, edge, padding, stride)
}
