package main

import (
	"fmt"
	"time"
)

func timeInUTC() time.Time {
	// Normally you'd fetch the time from a specific datasource here.
	return time.Date(2016, time.April, 26, 6, 10, 0, 0, time.UTC)
}

func main() {
	fmt.Println("Выводим текущее время в нулевом поясе:")
	x := time.Now().UTC()
	fmt.Println(x.Format("02-01-2006"))
	fmt.Println(x.Format("02-01-2006 15:04:05"))
	fmt.Println(x.Format(time.RFC822Z))

	fmt.Println("Возвращаем значение функции timeInUTC():")
	t := timeInUTC()
	fmt.Println(t)

	fmt.Println("Выводим время в соответствии с локацией пользователя:")
	userLocation, err := time.LoadLocation("Africa/Addis_Ababa") // "America/Los_Angeles" - время в NY, "Local" - местное время
	//список всех зон: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
	if err != nil {
		panic(err)
	}
	t = timeInUTC()        // присваиваем t - время в UTC 00
	t = t.In(userLocation) // присвиваем время в определенной локации ".In(локация)"
	fmt.Println(t.Format(time.RFC822Z))

	fmt.Println("Форматируемый вывод:")
	e, err := time.Parse("02-01-2006 15:04:05", "26-04-2016 08:10:23")
	if err != nil {
		panic(err)
	}
	fmt.Println(e.Format(time.RFC822Z))

	start := time.Now()
	for i := 0; i < 100; i++ {
	}
	duration := time.Since(start)
	fmt.Printf("Operation took: %s \n", duration)
}

// Домашнее задание. Что из статьи Time in Go: A primer стоит добавить в методичку:
// Считаю следует добавить информацию о том, как выводить время "UTC 00" и как выводить время в соответствии с локацией пользователя
// А также в методичку следует добавить информацию о том, как сравнивать даты а также как производить с ними операции
// Например func (Time) Sub, а также func (Time) Since
