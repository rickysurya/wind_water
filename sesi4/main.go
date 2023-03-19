package main

import (
	"fmt"
	"math"
)

// type biodata interface {
// 	getFullName() string

// }

// type mahasiswa struct {
// 	firstName, lastName string
// 	isMarried bool
// }

// func(m mahasiswa) getFullName() string {
// 	return m.firstName + " " + m.lastName
// }

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) perimeter() float64 {
	return math.Pi * (c.radius * 2)
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func main() {
	var c1 shape = circle{
		radius: 5,
	}
	var r1 shape = rectangle{
		width:  3,
		height: 2,
	}

	var c2 = circle{
		radius: 5,
	}
	fmt.Printf("Type of c1 : %T\n", c1)
	fmt.Printf("Type of r1 : %T\n", r1)
	fmt.Printf("Type of c2 : %T", c2)


}
