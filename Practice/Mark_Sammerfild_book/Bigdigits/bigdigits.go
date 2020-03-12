package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	var bigDigits = [][]string{
		{
			"  000  ",
			" 0   0 ",
			"0     0",
			"0     0",
			"0     0",
			" 0   0 ",
			"  000  "},
		{
			"   1 ",
			"  11 ",
			" 1 1 ",
			"   1 ",
			"   1 ",
			"   1 ",
			"  111",
		},
	}

	if len(os.Args) == 1 {
		fmt.Printf("usage: %s <whole-number>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	stringOfDigits := os.Args[1]

	// for row := 0; row < len(bigDigits[0]); row++ { //ряд
	// 	line := ""
	// 	for column := 0; column < len(stringOfDigits); column++ { //колонки
	// 		digit := stringOfDigits[column] - '0'
	// 		if 0 <= digit && digit <= 9 {
	// 			line += bigDigits[digit][row] + " "
	// 		} else {
	// 			log.Fatal("invalid whole number")
	// 		}
	// 	}
	// 	fmt.Println(line)
	// }

	for row := range bigDigits[0] { //ряд, строка
		line := ""
		for column := range stringOfDigits { //колонки
			digit := stringOfDigits[column] - '0'
			if 0 <= digit && digit <= 9 {
				line += bigDigits[digit][row] + " "
			} else {
				log.Fatal("invalid whole number")
			}
		}
		fmt.Println(line)
	}
}
