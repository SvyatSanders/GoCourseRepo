package main

import (
	"fmt"
	"math"
)

func main() {
	c := Circle{0, 0, 5}
	c2 := Circle{0, 0, 10}
	fmt.Println(c.area())
	fmt.Println(c2.area())

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())

	//fmt.Println(totalArea(&c, &r))

	fmt.Println(c.perimeter())
	fmt.Println(c2.perimeter())
	fmt.Println(r.perimeter())
}

// Shape - интерфейс для фигуры
type Shape interface {
	area() float64
	perimeter() float64
}

// Circle - круг
type Circle struct {
	x float64
	y float64
	r float64
}

// Rectangle - прямоугольник
type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

//Периметр круга равен произведению радиуса на два пи (3.1415).
func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2 * (l + w)
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

// MultiShape - Интерфейсы также могут быть использованы в качестве полей.
// Мы можем даже хранить в MultiShape данные Shape, определив в ней метод area
// Теперь MultiShape может содержать Circle, Rectangle и даже другие MultiShape.
type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}
