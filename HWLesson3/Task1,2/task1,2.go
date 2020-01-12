package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//1.	Описать несколько структур — любой легковой автомобиль и грузовик.
//Структуры должны содержать марку авто, год выпуска, объем багажника/кузова,
//запущен ли двигатель, открыты ли окна, насколько заполнен объем багажника.

//Car структура с описанием автомобилей
type Car struct {
	Brand             string
	Model             string
	EearOfIssue       int
	TrunkVolume       int
	EngineStarted     bool
	WindowsOpened     bool
	FilledTrunkVolume float64
}

//2.	Инициализировать несколько экземпляров структур.
//Применить к ним различные действия.
//Вывести значения свойств экземпляров в консоль.
func main() {
	//AddressBookMap()
	Toyota := Car{
		Brand:             "Toyota",
		Model:             "Supra",
		EearOfIssue:       2019,
		TrunkVolume:       290,
		EngineStarted:     true,
		WindowsOpened:     false,
		FilledTrunkVolume: 0.0,
	}

	Kamaz := Car{
		Brand:             "Kamaz",
		Model:             "65115",
		EearOfIssue:       2019,
		TrunkVolume:       6700,
		EngineStarted:     true,
		WindowsOpened:     false,
		FilledTrunkVolume: 0.0,
	}
	//вносим изменения в заданную структуру Toyota
	Toyota = Car{FilledTrunkVolume: 65.5}
	Toyota.EngineStarted = false

	fmt.Printf("\t %v \n %T \n", Toyota, &Toyota)
	fmt.Printf("\t %v \n %T \n", Toyota.FilledTrunkVolume, &Toyota.FilledTrunkVolume)
	fmt.Printf("\t %v \n\n", Kamaz)

	b, err := json.Marshal(Kamaz)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(b))
}
