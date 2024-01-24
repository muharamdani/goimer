# Goimer (Go Image Merger)

## Overview

Goimer is a package written in Go for merging a frame image with a set of second images. This allows you to create composite images by resizing and combining the frame and second images. It supports multiple image formats and provides flexibility in specifying the output directory and format.

## Features

- Merge a frame image with a directory of second images.
- Resize second images to match the dimensions of the frame image.
- Save the merged images in a specified output directory.
- Support for various image formats: JPEG, PNG, and WebP.

## Usage (CLI)
1. **Download the latest binary for your operating system from the [build](https://github.com/muharamdani/goimer/tree/master/build) directory.**
2. **Open a terminal and navigate to the directory containing the binary.**
3. **Run the binary with the following command:**
   - **Linux:** `./goimer_linux`
   - **Windows:** `goimer.exe`
   - **macOS:** `./goimer_darwin`
   - **Note:** You may need to grant the binary executable permissions. For example, on Linux, you can run `chmod +x goimer_linux` to grant executable permissions to the binary.
   - **Note:** According to the [Go FAQ](https://go.dev/doc/faq#virus), you may need to add an exception to your antivirus software to run the binary.
4. **Follow the prompts to enter the required information:**
   - Frame image filename
   - Directory containing the second images
   - Output directory
5. **Wait for the process to complete.**

## Usage (Package)
```bash
go get -u github.com/muharamdani/goimer
```
```go
package main

import "github.com/muharamdani/goimer/pkg/imagemerge"

func main() {
	// You can use fmt.Scanln() to get the input from user, or you can hardcode it
	cfg := imagemerge.MergeOptions{
		FrameImagePath: "frame/frame_file.png",
        SecondImageDir: "bunch_of_image_dir",
        OutputDir:      "output_dir",
	}
	
	if err := imagemerge.MergeImages(cfg); err != nil {
		return
	}
}
```
## Supported Image Formats

- JPEG
- PNG
- WebP

## Contributing

Contributions are welcome! If you find any issues or have ideas for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the [Apache License 2.0](LICENSE).

