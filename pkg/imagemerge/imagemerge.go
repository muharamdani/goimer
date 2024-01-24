package imagemerge

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/schollz/progressbar/v3"
	iDraw "golang.org/x/image/draw"
	"golang.org/x/image/webp"
)

const (
	jpgExt  = ".jpg"
	jpegExt = ".jpeg"
	pngExt  = ".png"
	webpExt = ".webp"
)

// MergeOptions represents the options for merging images.
type MergeOptions struct {
	FrameImagePath string
	SecondImageDir string
	OutputDir      string
}

// MergeImages merges a frame image with a set of second images.
func MergeImages(options MergeOptions) error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %v", err)
	}

	frameImgPath := filepath.Join(cwd, options.FrameImagePath)
	frameImg := loadImage(frameImgPath)

	secondImgDir := filepath.Join(cwd, options.SecondImageDir)
	outputDir := filepath.Join(cwd, options.OutputDir)

	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		if err := os.Mkdir(outputDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory %s: %v", outputDir, err)
		}
	}

	imageFiles, err := getImageFiles(secondImgDir)
	if err != nil {
		return fmt.Errorf("failed to get image files: %v", err)
	}

	var wg sync.WaitGroup

	bar := progressbar.Default(int64(len(imageFiles)), "Processing images")

	for _, imgPath := range imageFiles {
		wg.Add(1)
		go func(imgPath string) {
			defer wg.Done()

			secondImg := loadAndResizeImage(imgPath, frameImg.Bounds().Dx(), frameImg.Bounds().Dy())
			canvas := createCanvas(frameImg.Bounds())

			drawImageStretched(canvas, secondImg, frameImg.Bounds())
			drawImageOnTop(canvas, frameImg)

			outputFilePath := filepath.Join(outputDir, fmt.Sprintf("%s%s", strings.TrimSuffix(filepath.Base(imgPath), filepath.Ext(imgPath)), jpgExt))
			if err := saveImage(outputFilePath, canvas); err != nil {
				fmt.Printf("Error saving image %s: %v\n", outputFilePath, err)
			}

			bar.Add(1)
		}(imgPath)
	}

	wg.Wait()
	fmt.Println("All images processed.")

	return nil
}

func GetUserInput(prompt, cwd string) string {
	fmt.Print(prompt)
	var userInput string
	fmt.Scanln(&userInput)
	return filepath.Join(cwd, userInput)
}

func loadAndResizeImage(filePath string, width, height int) image.Image {
	file, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Errorf("failed to open file %s: %v", filePath, err))
	}
	defer file.Close()

	var img image.Image
	ext := strings.ToLower(filepath.Ext(filePath))

	switch ext {
	case jpgExt, jpegExt:
		img, err = jpeg.Decode(file)
	case pngExt:
		img, err = png.Decode(file)
	case webpExt:
		img, err = webp.Decode(file)
	default:
		err = fmt.Errorf("unsupported image format: %s", ext)
	}

	if err != nil {
		panic(fmt.Errorf("failed to decode image %s: %v", filePath, err))
	}

	return resizeImage(img, width, height)
}

func saveImage(filePath string, img image.Image) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %v", filePath, err)
	}
	defer file.Close()

	if err := jpeg.Encode(file, img, nil); err != nil {
		return fmt.Errorf("failed to encode image to JPEG: %v", err)
	}

	return nil
}

func getImageFiles(directory string) ([]string, error) {
	var imageFiles []string

	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if isImageFile(file.Name()) {
			imageFiles = append(imageFiles, filepath.Join(directory, file.Name()))
		}
	}

	return imageFiles, nil
}

func isImageFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".webp"
}

func loadImage(filePath string) image.Image {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	return img
}

func resizeImage(img image.Image, width, height int) image.Image {
	dstRect := image.Rect(0, 0, width, height)
	dst := image.NewRGBA(dstRect)
	iDraw.NearestNeighbor.Scale(dst, dstRect, img, img.Bounds(), draw.Over, nil)
	return dst
}

func createCanvas(bounds image.Rectangle) *image.RGBA {
	return image.NewRGBA(bounds)
}

func drawImageStretched(canvas draw.Image, img image.Image, bounds image.Rectangle) {
	imgRect := image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: bounds.Dx(), Y: bounds.Dy()},
	}
	draw.Draw(canvas, imgRect, img, image.Point{}, draw.Over)
}

func drawImageOnTop(canvas draw.Image, img image.Image) {
	draw.Draw(canvas, canvas.Bounds(), img, image.Point{}, draw.Over)
}
