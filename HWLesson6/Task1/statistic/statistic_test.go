package statistic

import "testing"

type testpair struct {
	values  []float64
	average float64
	summ    float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5, 3},
	{[]float64{1, 1, 1, 1, 1, 1}, 1, 6},
	{[]float64{-1, 6}, 2.5, 5},
	{[]float64{}, 0, 0},
}

func TestAverageSet(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}

// Дополните код из раздела «Тестирование» функцией подсчета суммы переданных элементов и тестом для этой функции.
func TestSumm(t *testing.T) {
	for _, pair := range tests {
		x := Summ(pair.values)
		if x != pair.summ {
			t.Error(
				"For", pair.values,
				"expected", pair.summ,
				"got", x,
			)
		}
	}
}
