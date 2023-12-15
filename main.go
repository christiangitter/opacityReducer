package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

func main() {

	// first we need to check the path of the file
	var inputFilePath string
	startDirectory := "assets/input"
	outputFileName := "assets/output/output.png"
	// we try to catch any errors that may occur during the walk
	err := filepath.Walk(startDirectory, func(path string, info os.FileInfo, err error) error {
		// if there is an error we return it
		if err != nil {
			return err
		}
		inputFilePath = path
		return nil
	})
	// if there is an error we log it
	if err != nil {
		log.Println(err)
	}

	// open the PNG file
	inputFile, err := os.Open(inputFilePath)

	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	// Decode the PNG image
	img, _, err := image.Decode(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new RGBA image to store the modified image
	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)

	// Loop through each pixel and reduce the opacity
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			oldColor := img.At(x, y)
			// Extract the existing alpha value
			r, g, b, oldAlpha := oldColor.RGBA()

			// Convert the alpha value to the range 0 to 255
			oldAlpha8 := uint8(oldAlpha >> 8)

			// Reduce the alpha value (you can modify this to change opacity)
			newAlpha8 := oldAlpha8 / 2

			// Convert the alpha value back to the range 0 to 65535
			newAlpha := uint16(newAlpha8) << 8

			// Create a new color with the modified alpha value
			newColor := color.NRGBA64{
				R: uint16(r),
				G: uint16(g),
				B: uint16(b),
				A: newAlpha,
			}

			// Set the pixel in the new image
			rgba.Set(x, y, newColor)
		}
	}

	// Save the modified image to a new PNG file
	outputFile, err := os.Create(outputFileName)

	fmt.Println("Output file: ", outputFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, rgba)
	if err != nil {
		log.Fatal(err)
	}

}
