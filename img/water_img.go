package main

//图片加水印

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

func main() {
	//原始图片是5.jpg
	imgb, _ := os.Open("5.jpg")
	img, _ := jpeg.Decode(imgb) //从imgb读取一幅jpeg格式的图像并解码返回该图像
	defer imgb.Close()

	wmb, _ := os.Open("test.png")
	watermark, _ := png.Decode(wmb) //从wmb读取一幅png格式的图像并解码返回该图像

	defer wmb.Close()

	//把水印写到右下角，并向0坐标各偏移10个像素  image.Pt返回Point{X , Y}  DX:返回r的宽度  Dy返回高度
	offset := image.Pt(img.Bounds().Dx()-watermark.Bounds().Dx()-10, img.Bounds().Dy()-watermark.Bounds().Dy()-10)
	b := img.Bounds()
	m := image.NewNRGBA(b) //NewNRGBA函数创建并返回一个具有指定范围的NRGBA

	//Draw方法通过使用Op参数调用包的Draw函数实现了Drawer接口
	draw.Draw(m, b, img, image.ZP, draw.Src) //ZP是原点
	draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)

	//生成新图片new.jpg，并设置图片质量..
	imgw, _ := os.Create("new.jpg")
	jpeg.Encode(imgw, m, &jpeg.Options{100})

	defer imgw.Close()

	fmt.Println("水印添加结束,请查看new.jpg图片...")
}
