package main

import "fmt"

type Point struct {
	x int
	y int
}

func NewPoint(y, x int) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p *Point) GetX() int {
	return p.x
}

func (p *Point) GetY() int {
	return p.y
}

func (p *Point) SetX(x int) {
	p.x = x
}

func (p *Point) SetY(y int) {
	p.y = y
}

func (p *Point) String() string {
	return fmt.Sprintf("{%d, %d}", p.y, p.x)
}

type Move struct {
	X int
	Y int
}

var availableMoves = []Move{
	Move{Y: 2, X: 1},
	Move{Y: 2, X: -1},
	Move{Y: 1, X: 2},
	Move{Y: -1, X: 2},
	Move{Y: 1, X: -2},
	Move{Y: -1, X: -2},
	Move{Y: -2, X: 1},
	Move{Y: -2, X: -1},
}

func CanItMoves(p Point, whereX, whereY int) bool {
	destX := p.x + whereX
	destY := p.y + whereY
	return (destX >= 1 && destX <= 8) && (destY >= 1 && destY <= 8)
}

func FindMove(p Point) []Point {
	result := make([]Point, 0, 2)
	for _, v := range availableMoves {
		if CanItMoves(p, v.X, v.Y) == true {
			result = append(result, NewPoint(p.GetY()+v.Y, p.GetX()+v.X))
		}
	}
	return result
}

func main() {
	p := NewPoint(0, 0)
	fmt.Println(FindMove(p))
	// resultS := make([]Point, 1, 2)
	// fmt.Printf("%v, %d, %d\n", resultS, len(resultS), cap(resultS))
}
