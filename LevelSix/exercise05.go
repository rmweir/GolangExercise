package main

import "fmt"

type square struct {
	length int
}

type circle struct {
	radius int
}

type shape interface {
	area() float64
}

func main() {
	sq := square{
		length: 2,
	}
	cir := circle{
		radius: 3,
	}
	display(sq)
	display(cir)
}

func display(s shape) {
	switch s.(type) {
	case square:
		fmt.Println("Shape is a square")
	case circle:
		fmt.Println("Shape is a circle")
	}
	fmt.Println(s.area())
}

func (s square) area() float64 {
	return float64(s.length * s.length)
}

func (c circle) area() float64 {
	return float64(c.radius*c.radius) * 3.14
}
