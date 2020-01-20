package statistic

import "math"

// Average - средний
func Average(xs []float64) float64 {
	var total float64
	for _, x := range xs {
		total += x
	}

	result := total / float64(len(xs))
	if math.IsNaN(result) {
		return 0
	}
	return result
}

// Дополните код из раздела «Тестирование» функцией подсчета суммы переданных элементов и тестом для этой функции.

// Summ - сумма переданных элементов
func Summ(xs []float64) float64 {
	var total float64
	for _, x := range xs {
		total += x
	}

	if math.IsNaN(total) {
		return 0
	}
	return total
}
