package main

import (
"math"
"fmt"
)

type geometry interface {

	Area() float64
	Perim() float64
}

type rect struct {

	length	float64
	width	float64
}

type square struct {

	side	float64
}

type circle struct {

	radius	float64
}

func (s square) Perim() float64 {

	return 4 * s.side
}

func (s square) Area() float64 {

	return math.Pow(s.side,2)
}

func (r rect) Area() float64 {

	return r.length * r.width
}

func (r rect) Perim() float64 {

	return 2.0 * (r.length + r.width)
}

func (c circle) Perim() float64 {

	return 2 * math.Pi * c.radius
}

func (c circle) Area() float64 {

	return math.Pi * math.Pow(c.radius, 2)
}

func measure(g geometry) {

	a := g.Area()
	p := g.Perim()

	fmt.Println(g, a, p)
}


func main() {


	r := rect{length:4, width: 2}
	c := circle{radius:5}
	s := square{side:9}

	a := r.Area()
	p := r.Perim()

	fmt.Println(r, a, p)
	measure(r)

	a = c.Area()
	p = c.Perim()

	fmt.Println(c, a, p)
	measure(c)

	a = s.Area()
	p = s.Perim()

	fmt.Println(s, a, p)
	measure(s)

}
