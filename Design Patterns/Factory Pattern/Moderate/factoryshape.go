package main

import (
	"fmt"
	"math"
)

type Shape interface {
	getArea() float64
}

type Circle struct {
	radius float64
}

func GetNewCircle(r float64) Shape {
	return &Circle{radius: r}
}

func (c *Circle) getArea() float64 {
	return math.Pi *c.radius*c.radius
}

type Square struct{
	side float64
}

func GetNewSquare(s float64) Shape {
	return &Square{side: s}
}

func (s *Square) getArea() float64{
	return s.side * s.side
}


func ShapeFactory(id int, s float64){
	var sh Shape
	if id == 1 {
		sh = GetNewCircle(s)
	} else {
		sh = GetNewSquare(s)
	}

	areaInfo := sh.getArea()
	fmt.Println(areaInfo)
}

func main(){
	ShapeFactory(1,1)
	ShapeFactory(2,2)
}