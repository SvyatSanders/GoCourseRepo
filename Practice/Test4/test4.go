package main

import "fmt"

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
func main() {
	fmt.Println(add(1, 2, 3))
	xs := []int{1, 2, 3, 4}
	fmt.Println(add(xs...))

	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())

	adds := func(x, y int) int {
		return x + y
	}
	fmt.Println(adds(10, -1))
}
