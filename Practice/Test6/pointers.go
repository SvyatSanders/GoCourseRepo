package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 1
}

func square(x *float64) {
	*x = *x * *x
}

func swap(aPtr *int, bPtr *int) {
	*aPtr, *bPtr = *bPtr, *aPtr
}

func main() {
	y := new(int)
	zero(y)
	fmt.Println(*y) // x всё еще равен 5
	fmt.Println(&y)

	x := 1.5
	square(&x)
	fmt.Println(x)

	a := 100
	b := 200
	swap(&a, &b)
	fmt.Printf("a=%v, b=%v\n", a, b)
}
