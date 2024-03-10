package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./path/mona.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	width := (img.Bounds().Size().X)
	height := img.Bounds().Size().Y
	aspectRatio := float64(width) / float64(height)

	// set the dimensions of the output ASCII art
	outputWidth := 144
	outputHeight := int(float64(outputWidth) / aspectRatio / 2)

	grayscale := make([][]uint8, height)
	for y := 0; y < height; y++ {
		grayscale[y] = make([]uint8, width)
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			luma := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			grayscale[y][x] = uint8(luma / 257) // scale to 0-255
		}
	}

	asciiChars := []string{" ", ".", ":", "-", "=", "+", "*", "#", "%", "@"}
	output := make([][]string, outputHeight)
	for y := 0; y < outputHeight; y++ {
		output[y] = make([]string, outputWidth)
		for x := 0; x < outputWidth; x++ {
			grayscaleX := int(float64(x) / float64(outputWidth) * float64(width))
			grayscaleY := int(float64(y) / float64(outputHeight) * float64(height))
			grayscaleValue := grayscale[grayscaleY][grayscaleX]
			asciiIndex := int(float64(grayscaleValue) / 25.5)
			output[y][x] = asciiChars[asciiIndex]
		}
		fmt.Println(strings.Join(output[y], "")) // print each row of ASCII art
	}
}
