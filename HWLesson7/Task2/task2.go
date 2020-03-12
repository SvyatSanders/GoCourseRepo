package main

import "fmt"

// 2.	Перепишите программу-конвейер, ограничив количество передаваемых для обработки значений
//  и обеспечив корректное завершение всех горутин.

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// генерация
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// возведение в квадрат
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	for {
		y, okk := <-squares
		if !okk {
			return
		}
		fmt.Println(y)
	}

}

// 0,1,4,9,16,25,36,49,64,81
