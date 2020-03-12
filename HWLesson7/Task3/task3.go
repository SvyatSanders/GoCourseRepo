package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

//3.	Добавьте для time-сервера возможность его завершения при вводе команды exit.
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("accepted connection from %v\n", conn.RemoteAddr())
		handleConn(conn) // горутина для многопоточного подключения к серверу
	}
}

func handleConn(c net.Conn) {
	cancel := make(chan int)
	go func() {
		var command string
		for {
			fmt.Fscan(os.Stdin, &command)
			if command == "exit" {
				cancel <- 1
				break
			}
			fmt.Println("Для выхода, ведите \"exit\"")
		}
	}()

	defer func() { //ToDo
		fmt.Printf("closing connection from %v\n", c.RemoteAddr())
		c.Close()
	}()

	for {
		tick := make(<-chan time.Time)
		tick = time.Tick(1 * time.Second)

		select {
		case <-tick:
			io.WriteString(c, time.Now().Format("15:04:05\n\r"))
		case <-cancel:
			fmt.Println("Connection canceled!")
			c.Close()
		}
	}
}
