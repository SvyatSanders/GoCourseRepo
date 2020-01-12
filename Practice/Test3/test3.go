package main

import "fmt"

func main() {
	x := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(x)
	fmt.Println(x[3])

	y := make([]int, 3, 9)
	fmt.Printf("%v %v\n", len(y), cap(y))

	fmt.Println(x[2:5]) //первые 2 значения слайса срезал, 3-5 значения вернул

	z := []int{48, 96, 86, 68, 57, 3, 63, 70, 37, 34, 83, 27, 19, 97, 8, 17}

	fmt.Println(smallestNumber(z))
}

func smallestNumber(s []int) int {
	smallest := s[0]
	for i := 0; i < len(s); i++ {
		if s[i] < smallest {
			smallest = s[i]
		}
	}
	return smallest
}
