package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

// Timeserver - создает сервер, который посылает клиенту каждую секунду текущее время
func main() {
	go spinner(20 * time.Millisecond) // вызывает функцию func spinner(delay time.Duration)
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn) // без go - последовательный, с go - параллельный !!!
	}
}

func spinner(delay time.Duration) { // крутится спиннер (-\|/)
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r) // c - символы
			time.Sleep(delay)
		}
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
