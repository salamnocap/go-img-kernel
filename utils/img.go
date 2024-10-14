package utils

import (
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"
)

func clamp(value float64, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func clampToByte(value float64) float64 {
	return clamp(value, 0, 255)
}

func LoadImage(filePath string) (image.Image, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func GetRGB(img image.Image) [][][]float64 {
	bounds := img.Bounds()

	rgbMatrix := make([][][]float64, bounds.Max.Y-bounds.Min.Y)
	for y := range rgbMatrix {
		rgbMatrix[y] = make([][]float64, bounds.Max.X-bounds.Min.X)
		for x := range rgbMatrix[y] {
			rgbMatrix[y][x] = make([]float64, 3)
		}
	}

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			r = r / 257
			g = g / 257
			b = b / 257
			rgbMatrix[y][x][0] = float64(r)
			rgbMatrix[y][x][1] = float64(g)
			rgbMatrix[y][x][2] = float64(b)
		}
	}

	return rgbMatrix
}

func GetGrayscale(img image.Image) [][]float64 {
	bounds := img.Bounds()

	grayMatrix := make([][]float64, bounds.Max.Y-bounds.Min.Y)
	for y := range grayMatrix {
		grayMatrix[y] = make([]float64, bounds.Max.X-bounds.Min.X)
	}

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gray := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b) // NTSC formula
			grayMatrix[y][x] = gray
		}
	}

	return grayMatrix
}

func LoadGrayImage(filePath string) ([][]float64, error) {
	img, err := LoadImage(filePath)
	if err != nil {
		return nil, err
	}
	return GetGrayscale(img), nil
}

func LoadRGBImage(filePath string) ([][][]float64, error) {
	img, err := LoadImage(filePath)
	if err != nil {
		return nil, err
	}
	return GetRGB(img), nil
}

func RgbToGray(rgbImage [][][]float64) [][]float64 {
	height := len(rgbImage)
	width := len(rgbImage[0])
	grayImage := make([][]float64, height)

	for y := 0; y < height; y++ {
		grayImage[y] = make([]float64, width)

		for x := 0; x < width; x++ {
			r := rgbImage[y][x][0]
			g := rgbImage[y][x][1]
			b := rgbImage[y][x][2]
			grayImage[y][x] = 0.299*r + 0.587*g + 0.114*b // NTSC formula
		}
	}

	return grayImage
}

func saveImageRGB(rgbImage [][][]float64, filePath string) error {
	height := len(rgbImage)
	if height == 0 {
		return fmt.Errorf("empty image")
	}
	width := len(rgbImage[0])
	if width == 0 {
		return fmt.Errorf("empty image")
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r := uint8(clampToByte(rgbImage[y][x][0]))
			g := uint8(clampToByte(rgbImage[y][x][1]))
			b := uint8(clampToByte(rgbImage[y][x][2]))
			img.Set(x, y, color.RGBA{R: r, G: g, B: b, A: 255})
		}
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return err
	}

	return nil
}

func saveImageGray(grayImage [][]float64, filePath string) error {
	height := len(grayImage)
	if height == 0 {
		return fmt.Errorf("empty image")
	}
	width := len(grayImage[0])
	if width == 0 {
		return fmt.Errorf("empty image")
	}

	img := image.NewGray(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			gray := uint8(clampToByte(grayImage[y][x]))
			img.SetGray(x, y, color.Gray{Y: gray})
		}
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return err
	}

	return nil
}

func SaveImage(img interface{}, path string, gray bool) error {
	var err error
	if gray {
		err = saveImageGray(img.([][]float64), path)
	} else {
		err = saveImageRGB(img.([][][]float64), path)
	}
	if err != nil {
		return err
	}
	fmt.Println("Image saved to", path)
	return nil
}
