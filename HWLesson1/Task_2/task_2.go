package main

import (
	"fmt"
	"math"
)

// 2.	Даны катеты прямоугольного треугольника.
// Найти его площадь, периметр и гипотенузу.
// Используйте тип данных float64 и функции из пакета math

func main() {
	var katet1 float64 = 2 // катет 1
	//fmt.Println("Введите длину 1-го катета")
	//fmt.Fscan(os.Stdin, &katet1)

	var katet2 float64 = 5 // катет 2
	//fmt.Println("Введите длину 2-го катета")
	//fmt.Fscan(os.Stdin, &katet2)

	var square = 0.5 * katet1 * katet2 // вычисляем площадь
	fmt.Println("Площадь прямоугольного треугольника = ", square, "см")

	hypotenuse := math.Sqrt(katet1*katet1 + katet2*katet2)
	//fmt.Printf("%.1f", hypotenuse)

	perimeter := katet1 + katet2 + hypotenuse
	//fmt.Println("Периметр прямоугольного треугольника = ", perimeter, "см")
	fmt.Println("Периметр прямоугольного треугольника = ", (Round(perimeter, 2)), "см")

	fmt.Println("Гипотенуза прямоугольного треугольника = ", (Round(hypotenuse, 2)), "см")
}

//Round - округление до 2х знаков после запятой
func Round(x float64, prec int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}
