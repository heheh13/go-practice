package main

import (
	"fmt"
	"math"
)

//interfaces can be used as types where ever i want to use
/// as a params type , return type , in  map ??

/// any shape having the area methods implemts the interfaces
type shape interface {
	area() float32
	//have() float32
}

type shape2 interface {
	haveFun() string
}

type circle struct {
	x, y, r float32
}

type rect struct {
	x, y float32
}

func (c *circle) area() float32 {
	return math.Pi * c.r * c.r
}

func (c circle) haveFun() string {

	return "having fun wiht circle"
}

func (r rect) area() float32 {
	return r.x * r.y
}

func (r rect) haveFun() string {
	return "having fun wiht rect"
}

func getArea(s shape) float32 {
	return s.area()
}

func main() {
	fmt.Println("interfaces")
	c := circle{0, 0, 5}
	r := rect{2, 3}

	shapes := []shape{&c, r}

	for _, shape := range shapes {
		fmt.Println(getArea(shape))
	}
	difShape := []shape2{c, r}
	for _, shape := range difShape {
		fmt.Println(shape.haveFun())
	}

}
