package squareroot

import (
	"fmt"
	"math"
)

// * Напишите функцию для вычисления корней квадратного уравнения (алгоритм можно найти в Википедии) и тесты к ней.
// ax2 + bx + c = 0

// SquareRoot - вычисляем корни квадратного уравнения
func SquareRoot(A, B, C float64) []float64 {
	a := A
	b := B
	c := C
	D := math.Pow(b, 2) - 4*a*c       //math.Pow(b, 2) - возведение в степень, где 2 - степерь, b - число
	fmt.Println("Дискриминант = ", D) //28
	var arr []float64
	if D > 0 {
		x1 := (-b + math.Sqrt(D)) / (2 * a)
		x2 := (-b - math.Sqrt(D)) / (2 * a)
		fmt.Printf("Так как дискриминант больше нуля то, квадратное уравнение имеет два действительных корня:\nx1 = %v\n", x1)
		fmt.Printf("x2 = %v\n", x2)
		arr = append(arr, x1, x2)
	} else if D == 0 {
		x := -b / (2 * a)
		fmt.Printf("Так как дискриминант больше нуля то, квадратное уравнение имеет один корень:\nx = %v\n", x)
		arr = append(arr, x)
	} else if D < 0 {
		fmt.Println("Так как дискриминант меньше нуля, то уравнение не имеет действительных решений.")
	}
	return arr
}

func main() {
	SquareRoot(2, -5, 3)
	SquareRoot(-4, 28, -49)
	SquareRoot(5, 6, 2)
}
