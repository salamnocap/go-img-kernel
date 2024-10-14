package transformation

import (
	"Go-Image-Kernels/kernel"
)

func boxBlurKernel() *kernel.Kernel {
	k := kernel.NewKernel(
		3,
		[][]float64{
			{1.0 / 9.0, 1.0 / 9.0, 1.0 / 9.0},
			{1.0 / 9.0, 1.0 / 9.0, 1.0 / 9.0},
			{1.0 / 9.0, 1.0 / 9.0, 1.0 / 9.0},
		},
	)
	return k
}

func BoxBlurRGB(rgbImage [][][]float64, padding int, stride int) [][][]float64 {
	blur := boxBlurKernel()
	return kernel.Convolve3D(rgbImage, blur, padding, stride)
}
