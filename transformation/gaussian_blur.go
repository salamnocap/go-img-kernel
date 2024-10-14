package transformation

import (
	"github.com/salamnocap/go-img-kernel/kernel"
)

func GaussianBlurKernel() *kernel.Kernel {
	k := kernel.NewKernel(
		3,
		[][]float64{
			{1.0 / 16, 2.0 / 16, 1.0 / 16},
			{2.0 / 16, 4.0 / 16, 2.0 / 16},
			{1.0 / 16, 2.0 / 16, 1.0 / 16},
		},
	)
	return k
}

func GaussianBlurKernel5() *kernel.Kernel {
	k := kernel.NewKernel(
		5,
		[][]float64{
			{1.0 / 256, 4.0 / 256, 6.0 / 256, 4.0 / 256, 1.0 / 256},
			{4.0 / 256, 16.0 / 256, 24.0 / 256, 16.0 / 256, 4.0 / 256},
			{6.0 / 256, 24.0 / 256, 36.0 / 256, 24.0 / 256, 6.0 / 256},
			{4.0 / 256, 16.0 / 256, 24.0 / 256, 16.0 / 256, 4.0 / 256},
			{1.0 / 256, 4.0 / 256, 6.0 / 256, 4.0 / 256, 1.0 / 256},
		},
	)
	return k
}

func GaussianBlurRGB(rgbImage [][][]float64, padding int, stride int) [][][]float64 {
	blur := GaussianBlurKernel()
	return kernel.Convolve3D(rgbImage, blur, padding, stride)
}

func GaussianBlurRGB5(rgbImage [][][]float64, padding int, stride int) [][][]float64 {
	blur := GaussianBlurKernel5()
	return kernel.Convolve3D(rgbImage, blur, padding, stride)
}
