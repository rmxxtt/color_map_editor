package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"math/rand"
)

func main() {
	fmt.Println("Генерация рандомной цветовой карты.")

	fmt.Print("Укажите количество цветов: ")
	var numberColors uint16
	_, err := fmt.Scanln(&numberColors)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Генерация %d цвета(ов) ...\n", numberColors)
	var colorMap = make([]color.RGBA, numberColors)
	generate(&colorMap)
	err = save(colorMap)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Готово!")
}

func generate(colorMap *[]color.RGBA) {
	for k := range *colorMap {
		(*colorMap)[k] = color.RGBA{
			R: uint8(rand.Intn(255)),
			G: uint8(rand.Intn(255)),
			B: uint8(rand.Intn(255)),
			A: uint8(rand.Intn(255)),
		}
	}
}

func save(colorMap []color.RGBA) error {
	data, err := json.MarshalIndent(colorMap, "", " ")
	if err != nil {
		return err
	}

	var path string
	fmt.Print("Укажите путь и название файла для сохранения (без формата): ")
	_, err = fmt.Scanln(&path)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path+".json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}
