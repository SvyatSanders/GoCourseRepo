package main

import (
	"fmt"
	"math"
	"os"
)

//1. Написать программу для конвертации рублей в доллары. Программа запрашивает сумму в рублях и выводит сумму в долларах. Курс доллара задайте константой.

//UsdCourse - курс доллара к рублю
const UsdCourse float64 = 63.29 // курс доллара

func main() {
	var deposit float64
	fmt.Println("Введите сумму вклада в рублях")
	fmt.Fscan(os.Stdin, &deposit)
	fmt.Println(Round((deposit/UsdCourse), 2), "USD")
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

// как правильно округлять
// как правильно проводить операции с разными типами данных (целочисленное и вещественное)
// exported function / const (type) should have comment or be unexported
