package main

//import "fmt"

//Внести в телефонный справочник дополнительные данные.
//Реализовать сохранение json-файла на диске с помощью пакета ioutil при изменении данных.

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	addressBook := make(map[string][]int) //Создаем адресную книгу

	addressBook["Alex"] = []int{89996543210}
	addressBook["Bob"] = []int{89167243812}
	addressBook["Bob"] = append(addressBook["Bob"], 89155243627)
	addressBook["Sarah"] = []int{87382230494}                        //Дополняем контактами
	addressBook["Sarah"] = append(addressBook["Sarah"], 89192343627) //Дополняем еще контактами

	fmt.Println(addressBook) //Печатаем адресную книгу как map

	for name, numbers := range addressBook { //Выводим в терминал адресную книгу форматированную
		fmt.Println("Абонент:", name)
		for i, number := range numbers {
			fmt.Printf("\t %v) %v \n", i+1, number)
		}
	}

	b, err := json.Marshal(addressBook) //Переводим адресную книгу из формата map в формат JSON
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("В файл result записан результат работы программы:\n")
	fmt.Println(string(b))

	WriteFile(string(b)) //Передаем функции WriteFile - данные справочника в формате JSON
}

//WriteFile - записываем в файл результат работы приложения
func WriteFile(text string) {
	toWrite := text
	message := []byte(toWrite)
	err := ioutil.WriteFile("/Users/SvyatSanders/go/src/GoCourseRepo/HWLesson3/Task4/result", message, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
