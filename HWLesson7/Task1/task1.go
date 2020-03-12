package main

// 1. Уберите из первого примера (Фибоначчи и спиннер) функцию, вычисляющую числа Фибоначчи.
// Как теперь заставить спиннер вращаться в течение некоторого времени? 10 секунд?

import (
	"fmt"
	"time"
)

// func main() {
// 	var wg sync.WaitGroup
// 	go spinner(20 * time.Millisecond)
// 	wg.Wait()

// 	timer1 := time.NewTimer(time.Second * 10) // Первый вариант с таймером
// 	<-timer1.C
// 	fmt.Println("Timer 1 expired")

// 	//time.Sleep(10 * time.Second)	// Второй вариант с Sleep
// }

// func spinner(delay time.Duration) {
// 	for {
// 		for _, r := range "-\\|/" {
// 			fmt.Printf("%c\r", r)
// 			time.Sleep(delay)
// 		}
// 	}

// }

func main() {
	go spinner(20 * time.Millisecond)

	timer1 := time.NewTimer(time.Second * 10) // Первый вариант с таймером
	<-timer1.C
	fmt.Println("Timer 1 expired")

	//time.Sleep(10 * time.Second) // Второй вариант с Sleep

	// c := make(chan bool) // Третий вариант с каналом
	// time.Sleep(time.Second * 2)
	// close(c)
}

func spinner(delay time.Duration) {

	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
			time.Sleep(delay)
		}
	}

}
