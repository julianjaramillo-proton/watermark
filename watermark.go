package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"strings"

	"github.com/fogleman/gg"
)

// userInput() reads the user input and returns it as a string
func userInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading input.", err)
	}
	input = strings.TrimSuffix(input, "\n")
	return input
}

// getImageFile() takes in a string containing the path to a desired image
// if the image is at the path given the function returns the image
func getImageFile(file string) image.Image {
	userImage, err := gg.LoadImage(file)
	if err != nil {
		log.Fatal("Error retrieving image file.", err)
	}
	return userImage
}

// watermark image
func watermark() {
	fmt.Println("Input image file to be watermarked")
	imagePath := userInput()

	inputImage := getImageFile(imagePath)

	fmt.Println("Enter text for watermark.")
	watermark := userInput()

	imageCtx := gg.NewContextForImage(inputImage)
	imageCtx.Rotate(gg.Radians(-10))
	imageCtx.SetColor(color.Alpha16{0x5fff})

	lenOfWatermark := len(watermark)
	for j := -100; j < imageCtx.Width()*2; j++ {
		for k := 0; k < imageCtx.Height()*2; k++ {
			if j%(lenOfWatermark*10) == 0 && k%20 == 0 {
				if k%40 == 0 {
					imageCtx.DrawStringAnchored(watermark, float64(j), float64(k), -1.5, .5)
				} else {
					imageCtx.DrawStringAnchored(watermark, float64(j), float64(k), .5, .5)
				}
			}
		}
	}

	fmt.Println("Enter desired name for watermarked file including .PNG/.png or .JPG/.jpg")
	fileName := userInput()
	fileType := typeOfImage(fileName)
	if fileType == 1 {
		saveAsJpg(fileName, imageCtx)
	} else if fileType == 2 {
		saveAsPng(fileName, imageCtx)
	} else {
		log.Fatal("Input image not jpg or png.")
	}
}

// saveAsJpg() saves the user specified watermarked image as a jpg/JPG
func saveAsJpg(fileName string, jpgImage *gg.Context) {
	jpgFile := fileName //+ ".JPG"
	filename, err := os.Create(jpgFile)
	if err != nil {
		panic(err)
	}
	defer filename.Close()
	if err = jpeg.Encode(filename, jpgImage.Image(), nil); err != nil {
		log.Printf("Encoding failure. %v", err)
	}
}

func saveAsPng(fileName string, pngImage *gg.Context) {
	png := fileName // + ".PNG"
	filename, err := os.Create(png)
	if err != nil {
		panic(err)
	}
	defer filename.Close()
	if err = jpeg.Encode(filename, pngImage.Image(), nil); err != nil {
		log.Printf("Encoding failure. %v", err)
	}
}

func typeOfImage(imgPath string) int {
	fileType := imgPath[len(imgPath)-3:]
	if fileType == "JPG" || fileType == "jpg" {
		return 1
	} else if fileType == "PNG" || fileType == "png" {
		return 2
	} else {
		log.Panic("Only JPG/jpg and PNG/png formats accepted")
		return 0
	}
}

func main() {
	watermark()
}
