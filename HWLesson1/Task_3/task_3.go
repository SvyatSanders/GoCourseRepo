package main

import (
	"fmt"
	"os"
)

// 3. Пользователь вводит сумму вклада в банк и годовой процент.
// Найти сумму вклада через 5 лет.

func main() {
	var vklad float64 // сумму вклада
	fmt.Println("Введите сумму вклада на 5 лет в рублях")
	fmt.Fscan(os.Stdin, &vklad)

	var percent float64 // сумму вклада
	fmt.Println("Введите ежегодный % вклада")
	fmt.Fscan(os.Stdin, &percent)

	firstYear := vklad * (1 + (percent / 100))
	secondYear := firstYear * (1 + (percent / 100))
	thirdYear := secondYear * (1 + (percent / 100))
	fourthYear := thirdYear * (1 + (percent / 100))
	fifthYear := fourthYear * (1 + (percent / 100))
	fmt.Printf("Сумма вклада через 5 лет = %.2f руб.\n", fifthYear)
}
