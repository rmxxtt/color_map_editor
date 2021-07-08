package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
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
	colorMapFile, err := ReadColorMapFile(path)
	if err != nil {
		fmt.Print(err)
		return
	}
	var colorMap []color.RGBA
	err = JsonUnmarshal(&colorMapFile, &colorMap)
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

func ReadColorMapFile(path string) ([]byte, error) {
	file, err := ioutil.ReadFile(path)
	return file, err
}

func ReadImageFile(path string, img *image.Image) error {
	file, err := os.Open(path)
	if err == nil {
		_ = file.Close()
	}

	*img, _, err = image.Decode(file)

	return err
}

func JsonUnmarshal(file *[]byte, v *[]color.RGBA) error {
	err := json.Unmarshal(*file, &v)
	return err
}
