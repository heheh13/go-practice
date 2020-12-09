package main

import (
	"fmt"
	"math"
)

//Doctor is doctore
//why i need commnet
type Doctor struct {
	number     int
	actorName  string
	companions []string
}

type animal struct {
	name   string
	origin string
}

type bird struct {
	animal
	speed  int32
	canFly bool
}

type point struct {
	x, y int
}

//Circle need doc
/// need doc
type Circle struct {
	x, y, r float32
}

func main() {

	aDoctor := Doctor{
		number:    3,
		actorName: "jhon partwee",
		companions: []string{
			"heheh hehe ",
			"jo grant",
			"shara jane",
		},
	}
	fmt.Println(aDoctor)

	aPerson := struct {
		name string
		age  int
	}{
		name: "jhon pertwee",
		age:  23,
	}
	fmt.Println(aPerson)
	b := bird{}
	b.name = "makyee"
	b.origin = "asia"
	b.speed = 65
	b.canFly = true

	b1 := bird{
		animal: animal{
			name:   "doyel",
			origin: "bd",
		},
		canFly: true,
		speed:  60,
	}

	fmt.Println(b)
	fmt.Println(b1)

	points := []point{}

	for i := 0; i < 10; i++ {
		p := point{x: i, y: i}
		points = append(points, p)
	}
	fmt.Println(points)

	c := new(Circle)
	*c = Circle{0, 0, 5}
	fmt.Println(circleArea(c))
	fmt.Println(c.area())

}
func circleArea(c *Circle) float32 {
	return math.Pi * c.r * c.r
}
func (c *Circle) area() float32 {
	return math.Pi * c.r * c.r
}
