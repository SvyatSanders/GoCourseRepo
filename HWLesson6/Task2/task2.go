package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

// Дополните пример из раздела «Пакет img» изображением горизонтальных и вертикальных линий.
// Воспользуйтесь статьей «Работа с изображениями».
// https://4gophers.ru/articles/rabota-s-izobrazheniyami/#.XiSJJlP7Q6g

var teal color.Color = color.RGBA{0, 200, 200, 255}
var red color.Color = color.RGBA{200, 30, 30, 255}

// Image - рисуем изображения
type Image interface {
	ColorModel() color.Model
	Bounds() image.Rectangle
	At(x, y int) color.Color
}

// Circle - круг
type Circle struct {
	p image.Point
	r int
}

// ColorModel - определяет какая цветовая модель будет использоваться для нашего изображения.
func (c *Circle) ColorModel() color.Model {
	return color.AlphaModel
}

// Bounds - определяем границы, в переделах которых будет размещена наша фигура.
func (c *Circle) Bounds() image.Rectangle {
	return image.Rect(c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r)
}

// At - возвращет цвет для определенных координат.
func (c *Circle) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.p.X)+0.5, float64(y-c.p.Y)+0.5, float64(c.r)
	if xx*xx+yy*yy < rr*rr {
		return color.Alpha{255}
	}
	return color.Alpha{0}
}

func main() {
	file, err := os.Create("someimage.png")

	if err != nil {
		fmt.Errorf("%s", err)
	}

	img := image.NewRGBA(image.Rect(0, 0, 500, 500))
	draw.Draw(img, img.Bounds(), &image.Uniform{teal}, image.ZP, draw.Src) //нарисовали бирюзовый фон
	// или draw.Draw(img, img.Bounds(), image.Transparent, image.ZP, draw.Src)

	for x := 20; x < 380; x++ { //нарисовали линию диагональную
		y := x/3 + 100
		img.Set(x, y, red)
	}

	for x := 20; x < 380; x++ { //нарисовали линию горизонтальную с отступом 20 от верхней границы (y)
		y := 20
		img.Set(x, y, red)
	}

	for x := 20; x < 380; x++ { //нарисовали линию горизонтальную с отступом 40 от верхней границы (y)
		y := 40
		img.Set(x, y, red)
	}

	for y := 20; y < 380; y++ { //нарисовали линию горизонтальную с отступом 40 от верхней границы (y)
		x := 20
		img.Set(x, y, red)
	}

	for y := 20; y < 380; y++ { //нарисовали линию горизонтальную с отступом 40 от верхней границы (y)
		x := 40
		img.Set(x, y, red)
	}

	mask := image.NewRGBA(image.Rect(0, 0, 500, 500))
	_ = mask
	//draw.Draw(mask, mask.Bounds(), &image.Uniform{red}, image.ZP, draw.Src)

	//draw.DrawMask(img, img.Bounds(), mask, image.ZP, &Circle{image.Point{20, 21}, 20}, image.ZP, draw.Over)
	//draw.DrawMask(img, img.Bounds(), mask, image.ZP, &Circle{image.Point{379, 141}, 20}, image.ZP, draw.Over)
	//draw.DrawMask(img, img.Bounds(), mask, image.ZP, &Circle{image.Point{239, 321}, 70}, image.ZP, draw.Over)

	png.Encode(file, img)
}
