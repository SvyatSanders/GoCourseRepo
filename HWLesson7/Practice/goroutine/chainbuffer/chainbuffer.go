package main

import "fmt"

func main() {
	fmt.Printf("%v \n", mirroredQuery())
}

func mirroredQuery() string {
	responses := make(chan string, 3) // responses - буферизированный канал, можно назвать очередью по fifo
	go func() {
		responses <- request("asia.site.io")
	}()
	go func() {
		responses <- request("americas.site.io")
	}()
	go func() {
		responses <- request("europe.site.io") //americas.site.io
	}()
	return <-responses // возврат самого быстрого ответа
}

func request(hostname string) string { /* */
	return hostname
}
