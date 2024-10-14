package test

import (
	"Go-Image-Kernels/transformation"
	"Go-Image-Kernels/utils"
	"testing"
)

var filePath = "../examples/input/castle_rgb.jpg"

func loadRGBImage(t *testing.T) [][][]float64 {
	img, err := utils.LoadRGBImage(filePath)
	if err != nil {
		t.Fatalf("Error loading image: %v", err)
	}
	return img
}

func saveImage(t *testing.T, img interface{}, path string, gray bool) {
	err := utils.SaveImage(img, path, gray)
	if err != nil {
		t.Fatalf("Error saving image: %v", err)
	}
}

func TestBinarization(t *testing.T) {
	rgbImage := loadRGBImage(t)
	binaryImage := transformation.BinarizationRGB(rgbImage, 128)
	saveImage(t, binaryImage, outputDir+"binary.jpg", true)
}

func TestBlur(t *testing.T) {
	rgbImage := loadRGBImage(t)
	blurredImage := transformation.GaussianBlurRGB(rgbImage, 1, 1)
	saveImage(t, blurredImage, outputDir+"blurred.jpg", false)
}

func TestEdgeDetection(t *testing.T) {
	rgbImage := loadRGBImage(t)
	edgeImage := transformation.EdgeDetectionRGB(rgbImage, 1, 1)
	saveImage(t, edgeImage, outputDir+"edge.jpg", true)
}

func TestDiagonalEdgeDetection(t *testing.T) {
	rgbImage := loadRGBImage(t)
	edgeImage := transformation.DiagonalEdgeDetectionRGB(rgbImage, 1, 1)
	saveImage(t, edgeImage, outputDir+"diagonal_edge.jpg", true)
}

func TestRidgeDetection(t *testing.T) {
	rgbImage := loadRGBImage(t)
	edgeImage := transformation.RidgeDetectionRGB(rgbImage, 1, 1)
	saveImage(t, edgeImage, outputDir+"ridge.jpg", true)
}

func TestGaussianBlur(t *testing.T) {
	rgbImage := loadRGBImage(t)
	blurredImage := transformation.GaussianBlurRGB(rgbImage, 1, 1)
	saveImage(t, blurredImage, outputDir+"gaussian_blurred.jpg", false)
}

func TestGaussianBlur5(t *testing.T) {
	rgbImage := loadRGBImage(t)
	blurredImage := transformation.GaussianBlurRGB5(rgbImage, 1, 1)
	saveImage(t, blurredImage, outputDir+"gaussian_blurred5.jpg", false)
}

func TestHorizontalTransformation(t *testing.T) {
	rgbImage := loadRGBImage(t)
	horizontalImage := transformation.HorizontalTransformationRGB(rgbImage, 1, 1)
	saveImage(t, horizontalImage, outputDir+"horizontal.jpg", false)
}

func TestIdentity(t *testing.T) {
	rgbImage := loadRGBImage(t)
	identityImage := transformation.IdentityTransformationRGB(rgbImage, 1, 1)
	saveImage(t, identityImage, outputDir+"identity.jpg", false)
}

func TestReflection(t *testing.T) {
	rgbImage := loadRGBImage(t)
	reflectionImage := transformation.ReflectionTransformationRGB(rgbImage, 1, 1)
	saveImage(t, reflectionImage, outputDir+"reflection.jpg", false)
}

func TestRotation(t *testing.T) {
	rgbImage := loadRGBImage(t)
	rotationImage := transformation.RotateRGBImage(rgbImage, 45)
	saveImage(t, rotationImage, outputDir+"rotation.jpg", false)
}

func TestSharpening(t *testing.T) {
	rgbImage := loadRGBImage(t)
	sharpenedImage := transformation.SharpenImageRGB(rgbImage, 1, 1)
	saveImage(t, sharpenedImage, outputDir+"sharpened.jpg", false)
}
