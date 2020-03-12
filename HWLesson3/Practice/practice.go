package main

import (
	"fmt"
)

func main() {
	//arr := [10]int{1, 2, 3, 4, 5}
	//fmt.Println(arr)

	slice := []int{1, 2, 3}
	fmt.Println(slice)
	fmt.Printf("slice = %T %v %v \n", slice, slice, cap(slice))

	slice = append(slice, 99, 99, 2, 3)
	//slice = append(slice, 99, 99, 2, 3)
	fmt.Printf("slice = %T %v %v \n", slice, slice, cap(slice))

	var slice3 = make([]int, len(slice))
	copy(slice3, slice)
	fmt.Printf("slice 3 = %T %v %v \n\n", slice3, slice3, cap(slice3))

	simpleMap()
}

func simpleMap() {
	var names = make(map[string]string)
	names["student1"] = "Alex"
	names["student2"] = "Stan"
	names["student3"] = "George"

	for student, n := range names {
		fmt.Println("Студент ", student)
		for i, name := range n {
			fmt.Printf("\t %v) %v\n", i+1, string(name))
		}
	}

	fmt.Printf("Студент №3 - %v", names["student4"])

	if _, ok := names["student4"]; !ok {
		fmt.Printf("Ключа student4 не существует \n\n")
	}
	// var a int
	// var b *int = &a
	// var c int = *b
	// fmt.Println(a, b, c)
}
