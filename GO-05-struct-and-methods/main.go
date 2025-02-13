package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	width, height, area float64
}

type circle struct {
	radius, area float64
}

type shape interface {
	getArea() float64
}

func calcArea(s shape) float64 {
	return s.getArea()
}

func (r *rectangle) getArea() float64 {
	return r.width * r.height
}

func (c *circle) getArea() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	r := rectangle{
		width:  25,
		height: 7,
	}
	c := circle{
		radius: 15.5,
	}
	r.area = calcArea(&r)
	c.area = calcArea(&c)

	fmt.Println("Area of the rectangle is: ", r.area)
	fmt.Println("Area of the circle is: ", c.area)

}
