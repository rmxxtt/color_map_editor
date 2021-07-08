package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"io/ioutil"
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
	var img image.Image
	err = ReadImageFile(path, &img)
	if err != nil {
		fmt.Print(err)
		return
	}
}

func ReadImageFile(path string, img *image.Image) error {
	file, err := os.Open(path)
	defer func(file *os.File) {
		e := file.Close()
		if e == nil {
			err = e
		}
	}(file)

	*img, _, err = image.Decode(file)

	return fmt.Errorf("")
}
