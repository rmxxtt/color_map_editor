package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/png"
	"io/ioutil"
	"math"
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
	file, err := os.Open(path)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer func(file *os.File) {
		e := file.Close()
		if e != nil {
			fmt.Print(err)
		}
	}(file)
	img, imgFormat, err := image.Decode(file)
	if err != nil {
		fmt.Print(err)
		return
	}

	newImg := ColorMapEditor(img, &colorMap)

	fmt.Print("Укажите путь сохранения изображения (без формата): ")
	_, err = fmt.Scanln(&path)
	if err != nil {
		fmt.Print(err)
		return
	}
	file, err = os.Create(path + "." + imgFormat)
	err = jpeg.Encode(file, newImg, nil)

	return
}

func ColorMapEditor(img image.Image, colorMap *[]color.RGBA) image.Image {
	if len(*colorMap) == 0 {
		return img
	}

	bounds := img.Bounds()
	size := img.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	newImg := image.NewRGBA(rect)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := RBGAtoUint8(img.At(x, y))
			newImg.Set(x, y, NewColor(c, colorMap))
		}
	}

	return newImg
}

func NewColor(c1 color.RGBA, colorMap *[]color.RGBA) color.RGBA {
	minDistance := math.MaxFloat64
	var r, g, b, a uint8

	for _, c2 := range *colorMap {
		distance := ColourDistance(c1, c2)
		if distance < minDistance {
			r, g, b, a = c2.R, c2.R, c2.R, c2.R
		}
	}

	return color.RGBA{R: r, G: g, B: b, A: a}
}

func ColourDistance(c1, c2 color.RGBA) (distance float64) {
	rMean := int32(c1.R+c2.R) / 2
	r := int32(c1.R) - int32(c2.R)
	g := int32(c1.G) - int32(c2.G)
	b := int32(c1.B) - int32(c2.B)
	distance = math.Sqrt(float64((((512 + rMean) * r * r) >> 8) + 4*g*g + (((767 - rMean) * b * b) >> 8)))

	return distance
}

func RBGAtoUint8(c color.Color) color.RGBA {
	R, G, B, A := c.RGBA()
	r := uint8(R >> 8)
	g := uint8(G >> 8)
	b := uint8(B >> 8)
	a := uint8(A >> 8)
	return color.RGBA{R: r, G: g, B: b, A: a}
}
