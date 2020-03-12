package main

import (
	"fmt"
	"math/big"
)

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

func fibMB() {
	// Initialize two big ints with the first two numbers in the sequence.
	a := big.NewInt(0)
	b := big.NewInt(1)

	// Initialize limit as 10^99, the smallest integer with 100 digits.
	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(99), nil)

	// Loop while a is smaller than 1e100.
	for a.Cmp(&limit) < 0 {
		// Compute the next Fibonacci number, storing it in a.
		a.Add(a, b)
		// Swap a and b so that b is the next number in the sequence.
		a, b = b, a
	}
	fmt.Println(a) // 100-digit Fibonacci number

	// Test a for primality.
	// (ProbablyPrimes' argument sets the number of Miller-Rabin
	// rounds to be performed. 20 is a good value.)
	fmt.Println(a.ProbablyPrime(20))
}

func main() {
	for i := 1; i <= 100; i++ {
		//fmt.Println(fibi(i))
		//fmt.Println(Fib(i))
	}
	fibMB()
}
