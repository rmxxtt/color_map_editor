package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
)

type Color struct {
	Red   uint8 `json:"red"`
	Green uint8 `json:"green"`
	Blue  uint8 `json:"blue"`
}

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

func generate(numberColors uint16) []Color {
	var colorMap = make([]Color, numberColors)
	for k := range colorMap {
		colorMap[k] = Color{Red: randomColor(), Green: randomColor(), Blue: randomColor()}
	}
	return colorMap
}

func randomColor() uint8 {
	min := 0
	max := 255
	return uint8(rand.Intn(max-min) + min)
}

func save(colorMap []Color) {
	data, err := json.MarshalIndent(colorMap, "", " ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("color_map.json", data, 0644)
	if err != nil {
		panic(err)
	}
}
