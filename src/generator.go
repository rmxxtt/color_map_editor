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

	var numberColors uint16
	fmt.Print("Укажите количество цветов: ")
	_, err := fmt.Scanln(&numberColors)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Генерация %d цвета(ов) ...\n", numberColors)
	colorMap := generate(numberColors)
	save(colorMap)
	fmt.Println("Готово!")
}

func generate(numberColors uint16) []color.RGBA {
	var colorMap = make([]color.RGBA, numberColors)
	for k := range colorMap {
		colorMap[k] = randomColor()
	}
	return colorMap
}

func randomColor() color.RGBA {
	return color.RGBA{
		R: randomUint8(),
		G: randomUint8(),
		B: randomUint8(),
		A: randomUint8(),
	}
}

func randomUint8() uint8 {
	return uint8(rand.Intn(255))
}

func save(colorMap []color.RGBA) {
	data, err := json.MarshalIndent(colorMap, "", " ")
	if err != nil {
		panic(err)
	}

	var path string
	fmt.Print("Укажите путь и название файла для сохранения (без формата): ")
	_, err = fmt.Scanln(&path)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(path+".json", data, 0644)
	if err != nil {
		panic(err)
	}
}
