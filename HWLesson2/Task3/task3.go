package main

import "fmt"

// 3. Написать функцию, которая последовательно выводит на экран 100 первых чисел Фибоначчи, начиная от 0.
// Числа Фибоначчи определяются соотношениями Fn=Fn-1 + Fn-2.

//Fib находим число Фибоначчи рекурсивно
func Fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return Fib(n-1) + Fib(n-2)
}

//Fib находим число Фибоначчи итеративно
func fibi(n int) int {
	var a, b int = 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(fibi(i))
		//fmt.Println(Fib(i))
	}
}

// a, b = b, a+b - непонятен пример
// просьба объяснить
