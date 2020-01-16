package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

// 3.	Самостоятельно изучите пакет encoding/csv. Напишите пример, иллюстрирующий его применение.

//Structure - структура, по которой сортируются данные из файла csv
type Structure struct {
	ID             string `json:"ID"`
	Sername        string `json:"Sername"`
	Name           string
	LastName       string
	Department     string
	Position       string
	Email          string
	MobileNumber   string
	CityNumber     string
	InternalNumber string
	HomeNumer      string
	Login          string
	Password       string
}

func main() {
	//csvFile, _ := os.Open("file.csv")

	bs, err := ioutil.ReadFile("test.csv")
	if err != nil {
		return
	}
	bss := string(bs)
	bsss := strings.ReplaceAll(bss, ";", ",")

	//fmt.Println(bsss)

	//reader := csv.NewReader(bufio.NewReader(csvFile))
	reader := csv.NewReader(strings.NewReader(bsss))
	var staff []Structure
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		staff = append(staff, Structure{
			ID:             line[0],
			Sername:        line[1],
			Name:           line[2],
			LastName:       line[3],
			Department:     line[4],
			Position:       line[5],
			Email:          line[6],
			MobileNumber:   line[7],
			CityNumber:     line[8],
			InternalNumber: line[9],
			HomeNumer:      line[10],
			Login:          line[11],
			Password:       line[12],
		},
		)
	}

	for name, numbers := range staff { //Выводим в терминал адресную книгу форматированную
		if name == 0 {
			fmt.Println("Названия полей:", name, numbers)
		} else {
			fmt.Println("Сотрудник:", name, numbers)
		}

		// for i, number := range numbers {
		// 	fmt.Printf("\t %v) %v \n", i+1, number)
		// }
	}

	//peopleJSON, _ := json.Marshal(staff)
	//fmt.Println(string(peopleJSON))
}

// не получилось вывести итоговую структуру красиво, как например, карту =(
