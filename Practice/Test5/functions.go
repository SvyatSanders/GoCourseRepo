package main

import "fmt"

func makeEvenGenerator() func() int {
	i := int(0)
	return func() (ret int) {
		ret = i
		i += 2
		return
	}
}

func main() {
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4

	arrayOfNum := []float64{2, 4}
	fmt.Println(sum(arrayOfNum))

	fmt.Println(half(2))

	fmt.Println(biggest(1, 2, 3, 44, 24))

	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd()) // 1
	fmt.Println(nextOdd()) // 3
	fmt.Println(nextOdd()) // 5
	fmt.Println(nextOdd()) // 7

	fmt.Println(fib(30))

	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}

func sum(arr []float64) float64 {
	var total float64
	for _, v := range arr {
		total += v
	}
	return total
}

func half(n int) (int, bool) {
	x := n / 2
	y := n % 2
	if y == 0 {
		return x, true
	}
	return x, false
}

func biggest(args ...int) int {
	biggest := args[0]
	for _, v := range args {
		if v > biggest {
			biggest = v
		}
	}
	return biggest
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func fib(n uint64) uint64 {
	return func() uint64 {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}()
}
