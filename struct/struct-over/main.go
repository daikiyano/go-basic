package main

import "fmt"

type Point struct {
	A int
	B string
	C float64
}

type BigPoint struct {
	Point Point
}

func main() {
	hp := BigPoint{}
	fmt.Println(hp)
	hp.Point.A = 100
	hp.Point.B = "Apple"
	hp.Point.C = 2.8

	fmt.Println(hp.A)

}
