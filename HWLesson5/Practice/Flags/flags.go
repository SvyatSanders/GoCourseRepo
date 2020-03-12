package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("str", "hello", "a string")
	numPtr := flag.Int("num", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var strVar string
	flag.StringVar(&strVar, "strVar", "world", "a string var")

	flag.Parse()

	fmt.Println("str:", *wordPtr)
	fmt.Println("num:", *numPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("strVar:", strVar)
	fmt.Println("tail:", flag.Args())
}

// go build task4.go
// ./task4 -str=hello -num=12 -fork -strVar=flag
// str: hello
// num: 12
// fork: true
// strVar: flag
// tail: []

// ./task4 -str=hello a1 a2 a3
// str: hello
// num: 42
// fork: false
// strVar: world
// tail: [a1 a2 a3]
