package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"os"
)

func main() {

	var path string
	fmt.Print("Укажите путь до изображения: ")
	//_, err := fmt.Scanln(&path)
	//if err != nil {
	//	fmt.Print(err)
	//	return
	//}
	path = "gen_test.jpg"
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
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Print(err)
		return
	}

	color1 := img.At(0, 0)
	fmt.Println("color 1\n", color1)

	r2, g2, b2, a2 := img.At(0, 0).RGBA()
	fmt.Println("r2, g2, b2, a2\n", r2, g2, b2, a2)
	fmt.Println("r2 >> 8, g2 >> 8, b2 >> 8, a2 >> 8\n", r2>>8, g2>>8, b2>>8, a2>>8)

	//color3 := img.At(0, 0).(color.RGBA)
	//fmt.Print("color3", color3)

	//_ = color.RGBAModel.Convert(img.At(0, 0).(color.RGBA))
	//fmt.Print("rgbacol\n", rgbacol)

	asd := image.NewRGBA(img.Bounds())
	fs := color.RGBAModel.Convert(asd.At(0, 0).(color.RGBA))
	fmt.Print("fs\n", fs)

	return
}
