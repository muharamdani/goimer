package main

import (
	"fmt"

	"github.com/muharamdani/goimer/pkg/imagemerge"
)

func main() {
	instruction := `
	Direction of use:
	1. Enter the frame image filename, e.g. frame.png, path/to/frame.png
	2. Enter the directory containing the second images, e.g. images, path/to/images
	3. Enter the output directory, e.g. output, path/to/output
	
	Note: Allowed image formats are png, jpg, jpeg, and webp`
	fmt.Println(instruction)

	frameImgPath := imagemerge.GetUserInput("Enter the frame image filename: ", "")
	secondImgDir := imagemerge.GetUserInput("Enter the directory containing the second images: ", "")
	outputDir := imagemerge.GetUserInput("Enter the output directory: ", "")

	options := imagemerge.MergeOptions{
		FrameImagePath: frameImgPath,
		SecondImageDir: secondImgDir,
		OutputDir:      outputDir,
	}

	if err := imagemerge.MergeImages(options); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
