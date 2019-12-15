package main

import (
	"fmt"
	"os"
)

//1. Написать программу для конвертации рублей в доллары. Программа запрашивает сумму в рублях и выводит сумму в долларах. Курс доллара задайте константой.

//UsdCourse - курс доллара к рублю
const UsdCourse float64 = 63.29 // курс доллара

func main() {
	var deposit float64
	fmt.Println("Введите сумму в рублях для конвертации в доллары")
	fmt.Fscan(os.Stdin, &deposit)
	fmt.Printf("рублей = %.2f USD\n", deposit/UsdCourse)
}
