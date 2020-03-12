package main

import (
	"flag"
	"io"
	"log"
	"os"
)

func main() {

	enabledFilePath := flag.String("file1", "test.txt", "enabled file Path")
	newFilePath := flag.String("file2", "test_copy.txt", "file path to copy")

	flag.Parse()

	x := *enabledFilePath
	y := *newFilePath

	sFile, err := os.Open(x)
	if err != nil {
		log.Fatal(err)
	}
	defer sFile.Close()

	eFile, err := os.Create(y)
	if err != nil {
		log.Fatal(err)
	}
	defer eFile.Close()

	_, err = io.Copy(eFile, sFile) // first var shows number of bytes
	if err != nil {
		log.Fatal(err)
	}

	err = eFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
}

// go build task4.go
// ./task4 -file1=test.txt -file2=test_copy2.txt - код в командной строке
// file1 - место текущего файла
// file2 - место куда копировать файл
//
// также программа работает и с абсолютным путем:
// ./task4 -file1=/Users/SvyatSanders/go/src/GoCourseRepo/HWLesson5/Task3/test.csv -file2=/Users/SvyatSanders/go/src/GoCourseRepo/HWLesson5/Task3/test2.csv
