package main

import (
	"fmt"
	"math"
)

//Разработать программу нахождения расстояния между двумя точками,
//которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

func main() {
	a := newPoint(12, 34)
	b := newPoint(-12, 15)
	fmt.Println(Distance(a, b))
}

type point struct {
	x, y float64
}

func newPoint(x, y float64) point {
	return point{x: x, y: y}
}

func Distance(p1, p2 point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}
