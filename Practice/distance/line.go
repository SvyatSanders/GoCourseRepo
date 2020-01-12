package distance

import "math"

//Line - линия между 2-мя точками
type Line struct {
	point1 Point
	point2 Point
}

//SetPointFirst - присваиваем значения первой точке с (X, Y)
func (l *Line) SetPointFirst(p Point) {
	l.point1 = p
}

//SetPointSecond - присваиваем значения второй точке с (X, Y)
func (l *Line) SetPointSecond(p Point) {
	l.point2 = p
}

//Distance - расстояние между 2-мя точками
func (l Line) Distance() (distance float64, err error) {
	distance = math.Sqrt(math.Pow(l.point2.X-l.point1.X, 2) + math.Pow(l.point2.Y-l.point1.Y, 2))
	return
}
