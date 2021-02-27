package main

import "fmt"

type Point struct {
	A int
	B string
	C float64
}

func Update(p Point) {
	p.A = 100
	p.B = "update"
	p.C = 2.14
}

// 参照型
func Update2(p *Point) {
	p.A = 100
	p.B = "update"
	p.C = 2.14
}

func main() {
	p := Point{}
	Update(p)
	fmt.Println(p)
	// 推奨されているやり方
	// アドレス演算子
	p2 := &Point{}
	Update2(p2)
	fmt.Println(p2)

	p3 := new(Point)
	Update2(p3)
	fmt.Println(p3)
}
