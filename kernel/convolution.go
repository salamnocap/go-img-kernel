package kernel

import "fmt"

func outputSize(width, height, padding, stride, kernelSize int) (int, int) {
	outputWidth := (width+2*padding-kernelSize)/stride + 1
	outputHeight := (height+2*padding-kernelSize)/stride + 1
	return outputWidth, outputHeight
}

func initializeOutput3D(height, width, channels int) [][][]float64 {
	output := make([][][]float64, height)
	for y := 0; y < height; y++ {
		output[y] = make([][]float64, width)
		for x := 0; x < width; x++ {
			output[y][x] = make([]float64, channels)
		}
	}
	return output
}

func initializeOutput2D(height, width int) [][]float64 {
	output := make([][]float64, height)
	for y := 0; y < height; y++ {
		output[y] = make([]float64, width)
	}
	return output
}

func applyPadding3D(input [][][]float64, padding int) [][][]float64 {
	if padding == 0 {
		return input
	}

	height := len(input)
	width := len(input[0])
	channels := len(input[0][0])

	paddedHeight := height + 2*padding
	paddedWidth := width + 2*padding

	padded := initializeOutput3D(paddedHeight, paddedWidth, channels)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			for c := 0; c < channels; c++ {
				padded[y+padding][x+padding][c] = input[y][x][c]
			}
		}
	}
	return padded
}

func applyPadding2D(input [][]float64, padding int) [][]float64 {
	if padding == 0 {
		return input
	}

	height := len(input)
	width := len(input[0])

	paddedHeight := height + 2*padding
	paddedWidth := width + 2*padding

	padded := initializeOutput2D(paddedHeight, paddedWidth)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			padded[y+padding][x+padding] = input[y][x]
		}
	}

	return padded
}

func Convolve3D(input [][][]float64, kernel *Kernel, padding int, stride int) [][][]float64 {
	height := len(input)
	width := len(input[0])
	channels := len(input[0][0])
	kernelSize := kernel.Size

	fmt.Printf(
		"🖼️ Input Image Size: %d pixels (Height: %d, Width: %d)\n",
		height*width, height, width,
	)
	fmt.Printf("🔍 Kernel Size: %d x %d\n", kernelSize, kernelSize)

	outputWidth, outputHeight := outputSize(width, height, padding, stride, kernelSize)
	output := initializeOutput3D(outputHeight, outputWidth, channels)
	paddedInput := applyPadding3D(input, padding)

	for y := 0; y < outputHeight; y++ {
		for x := 0; x < outputWidth; x++ {
			for c := 0; c < channels; c++ {
				sum := float64(0)
				for ky := 0; ky < kernelSize; ky++ {
					for kx := 0; kx < kernelSize; kx++ {
						inY := y*stride + ky
						inX := x*stride + kx
						sum += paddedInput[inY][inX][c] * kernel.Get(ky, kx)
					}
				}
				output[y][x][c] = sum
			}
		}
	}

	fmt.Printf(
		"🎨 Output Image Size: %d pixels (Height: %d, Width: %d)\n",
		len(output)*len(output[0]), len(output), len(output[0]),
	)
	return output
}

func Convolve2D(input [][]float64, kernel *Kernel, padding int, stride int) [][]float64 {
	height := len(input)
	width := len(input[0])
	kernelSize := kernel.Size

	fmt.Printf(
		"🖼️ Input Image Size: %d pixels (Height: %d, Width: %d)\n",
		height*width, height, width,
	)
	fmt.Printf("🔍 Kernel Size: %d x %d\n", kernelSize, kernelSize)

	outputWidth, outputHeight := outputSize(width, height, padding, stride, kernelSize)
	output := initializeOutput2D(outputHeight, outputWidth)
	paddedInput := applyPadding2D(input, padding)

	for y := 0; y < outputHeight; y++ {
		for x := 0; x < outputWidth; x++ {
			sum := float64(0)
			for ky := 0; ky < kernelSize; ky++ {
				for kx := 0; kx < kernelSize; kx++ {
					inY := y*stride + ky
					inX := x*stride + kx
					sum += paddedInput[inY][inX] * kernel.Get(ky, kx)
				}
			}
			output[y][x] = sum
		}
	}

	fmt.Printf(
		"🎨 Output Image Size: %d pixels (Height: %d, Width: %d)\n",
		len(output)*len(output[0]), len(output), len(output[0]),
	)
	return output
}
