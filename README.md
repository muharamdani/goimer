# Goimer (Go Image Merge Tool)

## Overview

Goimer is a command-line utility written in Go for merging a frame image with a set of second images. This tool allows you to create composite images by resizing and combining the frame and second images. It supports multiple image formats and provides flexibility in specifying the output directory and format.

## Features

- Merge a frame image with a directory of second images.
- Resize second images to match the dimensions of the frame image.
- Save the merged images in a specified output directory.
- Support for various image formats: JPEG, PNG, and WebP.

## Usage

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/muharamdani/goimer.git
   ```

2. **Build:**

   ```bash
   cd goimer
   ./build.sh
   ```

3. **Run:**
   
   ```bash
   # For Linux   
   ./build/goimer_linux
   
   # For windows
   run file from ./build/goimer.exe
   ```

4. **Follow the prompts to enter the required information:**
    - Frame image filename
    - Directory containing the second images
    - Output directory

5. **Wait for the process to complete.**

## Supported Image Formats

- JPEG
- PNG
- WebP

## Contributing

Contributions are welcome! If you find any issues or have ideas for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the [Apache License 2.0](LICENSE).

