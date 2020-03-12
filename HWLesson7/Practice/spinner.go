package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(20 * time.Millisecond) // вызывает функцию func spinner(delay time.Duration)
	const n = 43
	fibN := fibonacci(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) { // крутится спиннер (-\|/)
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r) // c - символы
			time.Sleep(delay)
		}
	}
}

func fibonacci(x int) int {
	if x < 2 {
		return x
	}
	return fibonacci(x-1) + fibonacci(x-2)
}
