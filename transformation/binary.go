package transformation

import "github.com/salamnocap/go-img-kernel/utils"

func BinarizationGray(grayImage [][]float64, threshold float64) [][]float64 {
	for i := 0; i < len(grayImage); i++ {
		for j := 0; j < len(grayImage[i]); j++ {
			if grayImage[i][j] >= threshold {
				grayImage[i][j] = 255
			} else {
				grayImage[i][j] = 0
			}
		}
	}
	return grayImage
}

func BinarizationRGB(rgbImage [][][]float64, threshold float64) [][]float64 {
	grayImage := utils.RgbToGray(rgbImage)
	return BinarizationGray(grayImage, threshold)
}
