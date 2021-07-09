package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/png"
	"io/ioutil"
	"math/rand"
	"os"
)

func main() {
	fmt.Println("Редактор цветовой палитры изображения.")

	fmt.Print("Укажите путь до файла Цветовой карты: ")
	var path string
	_, err := fmt.Scanln(&path)
	if err != nil {
		fmt.Print(err)
		return
	}
	colorMapFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
		return
	}
	var colorMap []color.RGBA
	err = json.Unmarshal(colorMapFile, &colorMap)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print("Укажите путь до изображения: ")
	_, err = fmt.Scanln(&path)
	if err != nil {
		fmt.Print(err)
		return
	}
	img, imgFormat, err := ReadImageFile(path)
	if err != nil {
		fmt.Print(err)
		return
	}

	err = pass(img, imgFormat, colorMap)
	if err != nil {
		fmt.Print(err)
		return
	}
}

func ReadImageFile(path string) (image.Image, string, error) {
	file, err := os.Open(path)
	defer func(file *os.File) {
		e := file.Close()
		if err == nil {
			err = e
		}
	}(file)

	img, imgFormat, err := image.Decode(file)

	return img, imgFormat, err
}

func pass(img image.Image, imgFormat string, colorMap []color.RGBA) error {
	bounds := img.Bounds()
	size := img.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	newImg := image.NewRGBA(rect)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			r, g, b, a = r>>8, g>>8, b>>8, a>>8
			newImg.Set(x, y, color.RGBA{
				R: uint8(rand.Intn(255)),
				G: uint8(rand.Intn(255)),
				B: uint8(rand.Intn(255)),
				A: uint8(rand.Intn(255)),
			})
		}
	}

	fmt.Print("Укажите путь сохранения изображения (без формата): ")
	var path string
	_, err := fmt.Scanln(&path)

	file, err := os.Create(path + "." + imgFormat)
	err = jpeg.Encode(file, newImg, nil)

	return err
}
