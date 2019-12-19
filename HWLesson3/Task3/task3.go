//FIFO Реализовать очередь.
//Это структура данных, работающая по принципу FIFO
//(First Input — first output, или «первым зашел — первым вышел»)

//FIFO in Golang
package main

import "fmt"

func main() {
	var queue []string

	queue = append(queue, "1 element") // Добавляем элементы
	queue = append(queue, "2 element")
	queue = append(queue, "3 element", "4 element", "5 element", "6 element")

	fmt.Printf("Слайс состоит из %v \n", queue)

	// Когда стек будет пустым, он вернет пустую строку
	for len(queue) >= 0 {
		if len(queue) > 0 {
			fmt.Printf("Уаляем %v \n", queue[0]) // Печатаем первый элемент
			queue = queue[1:]                    // Обрезаем слайс от первого элемента [:0]
		} else if len(queue) == 0 {
			fmt.Println("Удалили все элементы")
			break
		}
	}
}
