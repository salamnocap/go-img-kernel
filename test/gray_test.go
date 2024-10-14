package test

import (
	"github.com/salamnocap/go-img-kernel/transformation"
	"github.com/salamnocap/go-img-kernel/utils"
	"testing"
)

var grayFilePath = "../examples/input/castle_gray.jpg"
var outputDir = "../examples/output/"

func loadGrayImage(t *testing.T) [][]float64 {
	img, err := utils.LoadGrayImage(grayFilePath)
	if err != nil {
		t.Fatalf("Error loading grayscale image: %v", err)
	}
	return img
}

func saveGrayImage(t *testing.T, img interface{}, path string) {
	err := utils.SaveImage(img, path, true)
	if err != nil {
		t.Fatalf("Error saving grayscale image: %v", err)
	}
}

func TestBinarizationGray(t *testing.T) {
	grayImage := loadGrayImage(t)
	binaryImage := transformation.BinarizationGray(grayImage, 0.5)
	saveGrayImage(t, binaryImage, outputDir+"binary_gray.jpg")
}

func TestEdgeDetectionGray(t *testing.T) {
	grayImage := loadGrayImage(t)
	edgeImage := transformation.EdgeDetectionGray(grayImage, 1, 1)
	saveGrayImage(t, edgeImage, outputDir+"edge_gray.jpg")
}

func TestDiagonalEdgeDetectionGray(t *testing.T) {
	grayImage := loadGrayImage(t)
	edgeImage := transformation.DiagonalEdgeDetectionGray(grayImage, 1, 1)
	saveGrayImage(t, edgeImage, outputDir+"diagonal_edge_gray.jpg")
}

func TestHorizontalTransformationGray(t *testing.T) {
	grayImage := loadGrayImage(t)
	horizontalImage := transformation.HorizontalTransformationGray(grayImage, 1, 1)
	saveGrayImage(t, horizontalImage, outputDir+"horizontal_gray.jpg")
}

func TestIdentityGray(t *testing.T) {
	grayImage := loadGrayImage(t)
	identityImage := transformation.IdentityTransformationGray(grayImage, 1, 1)
	saveGrayImage(t, identityImage, outputDir+"identity_gray.jpg")
}

func TestReflectionGray(t *testing.T) {
	grayImage := loadGrayImage(t)
	reflectionImage := transformation.ReflectionTransformationGray(grayImage, 1, 1)
	saveGrayImage(t, reflectionImage, outputDir+"reflection_gray.jpg")
}

func TestSharpeningGray(t *testing.T) {
	grayImage := loadGrayImage(t)
	sharpenedImage := transformation.SharpenImageGray(grayImage, 1, 1)
	saveGrayImage(t, sharpenedImage, outputDir+"sharpened_gray.jpg")
}
