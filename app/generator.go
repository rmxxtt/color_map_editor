package main

import (
	"fmt"
	"math/rand"
)

type Color struct {
	red   uint8
	green uint8
	blue  uint8
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
}

func generate(numberColors uint16) []Color {
	var colorMap = make([]Color, numberColors)
	for k := range colorMap {
		colorMap[k] = Color{red: randomColor(), green: randomColor(), blue: randomColor()}
	}
	return colorMap
}

func randomColor() uint8 {
	min := 0
	max := 200
	return uint8(rand.Intn(max-min) + min)
}

func save(colorMap []Color) {
	println(colorMap)
}
