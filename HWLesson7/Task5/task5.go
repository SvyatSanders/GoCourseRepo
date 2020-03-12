package main

// Попробуйте смоделировать с помощью sync.WaitGroup гонку автомобилей.
// Каждый автомобиль должен быть представлен отдельной горутиной со случайно устанавливаемой скоростью (math/rand)
// и случайным временем готовности к старту.
// Программа должна ожидать готовности всех машин, обеспечивать одновременный старт и фиксировать финиш каждой машины.

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	raceDuration = 5 * time.Second
)

//Car - машина
type Car struct {
	Name        string
	Speed       int
	PrepareTime time.Duration

	position int
}

//Prepare - готовность
func (c *Car) Prepare(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(c.PrepareTime)
	fmt.Printf("Car %q is ready\n", c.Name)
}

//GetPosition - получение позиции
func (c *Car) GetPosition() int {
	return c.position
}

//Race - гонка
func (c *Car) Race(raceDuration time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	end := time.After(raceDuration)
	for range ticker.C {
		select {
		case <-end:
			return
		default:
		}
		c.position += c.Speed
		fmt.Printf("Car %q has position %d\n", c.Name, c.position)
	}
}

func findWinner(cars ...Car) (Car, error) {
	maxPosition := -1
	maxIdx := -1
	for i, v := range cars {
		if v.position >= maxPosition {
			maxPosition = v.position
			maxIdx = i
		}
	}

	if maxIdx == -1 {
		return Car{}, errors.New("no winner found")
	}

	return cars[maxIdx], nil
}

func main() {
	car1 := Car{
		Name:        "Lada 4x4",
		Speed:       1,
		PrepareTime: 200 * time.Millisecond,
	}
	car2 := Car{
		Name:        "Lada Semerka",
		Speed:       2,
		PrepareTime: 100 * time.Millisecond,
	}
	car3 := Car{
		Name:        "Daewoo Nexia",
		Speed:       4,
		PrepareTime: 300 * time.Millisecond,
	}

	var wg sync.WaitGroup
	wg.Add(3)
	go car1.Prepare(&wg)
	go car2.Prepare(&wg)
	go car3.Prepare(&wg)

	wg.Wait()
	fmt.Println("Starting the race!")

	wg.Add(3)
	go car1.Race(raceDuration, &wg)
	go car2.Race(raceDuration, &wg)
	go car3.Race(raceDuration, &wg)
	wg.Wait()
	fmt.Println("The race has ended!")

	winner, err := findWinner(car1, car2, car3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("The winner is %q", winner.Name)
}
