package main

import ("fmt"
"math"
)

type Shape interface{
	area() float64
}
type DummyShape interface {
	area() float64
}
type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	length, breadth float64
}

func (c Circle) area() float64{
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.length * r.breadth
}

func getArea(s Shape) float64 {
	return s.area()
}

func getDummyArea(s DummyShape) float64 {
	return s.area()
}
func main() {
	circle := Circle{0, 0, 5}
	rectangle := Rectangle{10, 10}
	fmt.Println("Circle's area is ", getArea(circle))
	fmt.Println("Rectangle's area is ", getArea(rectangle))
	fmt.Println("Dummy area is ", getDummyArea(rectangle))
}

