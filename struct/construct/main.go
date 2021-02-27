package main

import "fmt"

// Point 構造体
type Point struct {
	A int
	B string
	C float64
}

// NewPoint 構造体
func NewPoint(a int, b string, c float64) *Point {
	return &Point{a, b, c}
}

func main() {
	p1 := Point{1, "A", 1.1}
	p2 := NewPoint{1, "A", 1.1}
	fmt.Println(p1)
	fmt.Println(p2)
}
