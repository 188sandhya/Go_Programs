package main

import (
	"fmt"
	"math"
)

type geometry interface {
	Area() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) Area() float64 {
	return r.height * r.width
}

func (c circle) Area() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println("g:", g)
	fmt.Println(g.Area())

}

func main() {
	re := rect{
		width:  10,
		height: 10,
	}
	fmt.Println("rect Area", re.Area())

	r := rect{3, 4}
	c := circle{5}

	measure(r)
	measure(c)

}
