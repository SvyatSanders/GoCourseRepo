package main

import (
	"fmt"
)

func main() {
	ourSlice := []int{98, 93, 77, 82, 83}

	//Получаем все аргументы
	res, count := average(10, 20, 14, 100)
	fmt.Println(res)
	fmt.Println(count)

	//Игнорируем все аргументы
	average(ourSlice...)

	//Игнорируем только второй аргумент
	res, _ = average(13, 24, 93, 12, 43)
}

func average(values ...int) (res int, count int) {
	count = len(values)
	total := 0
	for _, v := range values {
		total += v
	}
	res = total / len(values)
	return
}
