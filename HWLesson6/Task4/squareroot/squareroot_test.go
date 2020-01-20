package squareroot

import (
	"testing"
)

type testpair struct {
	A float64
	B float64
	C float64
	D float64
	X []float64
}

var arr []float64 // - пустой слайс (nil) для проверки при "D < 0"
var values = []testpair{
	{2, -5, 3, 1, []float64{1.5, 1}},
	{-4, 28, -49, 0, []float64{3.5}},
	{5, 6, 2, -4, arr},
}

func TestSquareRoot(t *testing.T) {
	for _, pair := range values {
		v := SquareRoot(pair.A, pair.B, pair.C)
		if compare(v, pair.X) != true {
			t.Error(
				"For", pair.A, pair.B, pair.C,
				"expected", pair.X,
				"got", v,
			)
		}
	}
}

func compare(a, b []float64) bool {
	if len(a) == len(b) {
		if len(a) == 2 {
			if a[0] == b[0] && a[1] == b[1] {
				return true
			}
			return false
		} else if len(a) == 1 {
			if a[0] == b[0] {
				return true
			}
			return false
		} else if len(a) == 0 {
			if a == nil && b == nil {
				return true
			}
			return false
		}
		return false
	}
	return false
}
