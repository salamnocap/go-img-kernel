package transformation

import (
	"math"
)

func RotateRGBImage(rgbImage [][][]float64, degree int) [][][]float64 {
	height := len(rgbImage)
	width := len(rgbImage[0])

	theta := float64(degree) * math.Pi / 180.0
	cx := float64(width) / 2.0
	cy := float64(height) / 2.0

	rotatedImage := make([][][]float64, height)
	for y := 0; y < height; y++ {
		rotatedImage[y] = make([][]float64, width)
		for x := 0; x < width; x++ {
			rotatedImage[y][x] = make([]float64, 3)
		}
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			xf := float64(x) - cx
			yf := float64(y) - cy
			srcX := int(cx + xf*math.Cos(theta) + yf*math.Sin(theta))
			srcY := int(cy - xf*math.Sin(theta) + yf*math.Cos(theta))

			if srcX >= 0 && srcX < width && srcY >= 0 && srcY < height {
				rotatedImage[y][x][0] = rgbImage[srcY][srcX][0]
				rotatedImage[y][x][1] = rgbImage[srcY][srcX][1]
				rotatedImage[y][x][2] = rgbImage[srcY][srcX][2]
			} else {
				rotatedImage[y][x][0] = 255
				rotatedImage[y][x][1] = 255
				rotatedImage[y][x][2] = 255
			}
		}
	}

	return rotatedImage
}
