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
	fmt.Printf("Периметр прямоугольного треугольника = %.2f см\n", perimeter)

	fmt.Printf("Гипотенуза прямоугольного треугольника = %.2f см\n", hypotenuse)
}
