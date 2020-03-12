package main

import "fmt"

//1. Написать свой интерфейс и создать несколько структур, удовлетворяющих ему.

//Vehicle - транспортное средство
type Vehicle interface {
	Destination() //метод
}

//Car - структура "Автомобиль"
type Car struct {
}

//Aircraft - структура "Самолет"
type Aircraft struct {
}

//Destination - метод, описывающий что делает автомобиль
func (c Car) Destination() {
	fmt.Println("Машина едет")
}

//Destination - метод, описывающий что делает автомобиль
func (a Aircraft) Destination() {
	fmt.Println("Самолет летит")
}

func drive(v Vehicle) {
	v.Destination()
}

func main() {
	tesla := Car{}
	boing := Aircraft{}
	drive(tesla)
	drive(boing)
}
