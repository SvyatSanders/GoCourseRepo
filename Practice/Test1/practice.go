package main

import (
	"GoCourseRepo/Practice/distance"
	"fmt"
	"math"
)

func main() {
	var c Circle
	c = Circle{x: 0, y: 0, r: 5}

	fmt.Println(c.x, c.y, c.r)
	c.x = 10
	c.y = 5

	fmt.Printf("Площадь круга: %0.2f\n", c.area())

	liiine := distance.Line{}
	liiine.SetPointFirst(distance.Point{1, 0})
	liiine.SetPointSecond(distance.Point{0, 1})

	fmt.Printf("Массив Line из пакета distance: %v\n", liiine)

	lineDistance, _ := liiine.Distance()
	fmt.Printf("Расстояние между 2-мя точками: %0.2f\n", lineDistance)
}

//Circle - круг
type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}
