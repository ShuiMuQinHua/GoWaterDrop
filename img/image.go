package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
)

func main() {
	f1, err := os.Open("1.jpg")
	if err != nil {
		panic(err)
	}
	defer f1.Close()
	f2, err := os.Open("2.jpg")
	if err != nil {
		panic(err)
	}
	defer f2.Close()
	f3, err := os.Create("3.jpg")
	if err != nil {
		panic(err)
	}
	defer f3.Close()

	m1, err := jpeg.Decode(f1) //从f1读取一副jpeg格式的图像，并解码返回该图像
	if err != nil {
		panic(err)
	}
	bounds := m1.Bounds() //该函数返回一个矩形(返回f1的矩形)

	m2, err := jpeg.Decode(f2) //从f2读取一副jpeg格式的图像，并解码返回该图像
	if err != nil {
		panic(err)
	}

	m := image.NewRGBA(bounds) //NewRGBA函数创建并返回一个具有指定范围的RGBA(RGBA类型代表一幅内存中的图像)

	white := color.RGBA{255, 255, 255, 255}                         //?
	draw.Draw(m, bounds, &image.Uniform{white}, image.ZP, draw.Src) //Draw函数使用nil的mask参数调用DrawMask函数。
	draw.Draw(m, bounds, m1, image.ZP, draw.Src)
	draw.Draw(m, image.Rect(25, 50, 75, 150), m2, image.Pt(200, 60), draw.Src)
	err = jpeg.Encode(f3, m, &jpeg.Options{90})
	if err != nil {
		panic(err)
	}
	fmt.Printf("ok\n")

}
