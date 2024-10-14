package transformation

import (
	"github.com/salamnocap/go-img-kernel/kernel"
)

func BoxBlurKernel() *kernel.Kernel {
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
	blur := BoxBlurKernel()
	return kernel.Convolve3D(rgbImage, blur, padding, stride)
}
