package main

import (
	"fmt"
	"os"
)

// 1. Написать функцию, которая определяет, четное ли число.
func main() {
	var number int
	fmt.Println("Введите число для проверки числа на четность: ")
	fmt.Fscan(os.Stdin, &number)
	var check = number % 2
	if check == 0 {
		fmt.Println("Введенное число ", number, " является четным")
	} else {
		fmt.Println("Введенное число ", number, " является НЕчетным")
	}
}
