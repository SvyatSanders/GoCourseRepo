package main

import (
	"fmt"
	"sort"
)

// 2. Создать тип, описывающий контакт в телефонной книге.
// Создать псевдоним типа телефонной книги (массив контактов) и реализовать для него
// интерфейс Sort{}.

//Person (записи) - структура, описывающая телефонную книгу
type Person struct {
	Name        string
	Age         int
	PhoneNumber []int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d, %d", p.Name, p.Age, p.PhoneNumber)
}

// ByAge implements sort.Interface for []Person based on the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age > a[j].Age }

func main() {
	adressBook := []Person{
		{"Миша", 38, []int{78293467382}},
		{"Никита", 25, []int{89167253764, 89635437382}},
		{"Ирина", 42, []int{89267252344, 89265432488}},
		{"Сара", 31, []int{87382230494, 89192343627, 89192343623}},
	}

	fmt.Println(adressBook)
	sort.Sort(ByAge(adressBook))
	fmt.Println(adressBook)

	sort.Slice(adressBook, func(i, j int) bool {
		return adressBook[i].Age > adressBook[j].Age
	})
	fmt.Println(adressBook)
}
