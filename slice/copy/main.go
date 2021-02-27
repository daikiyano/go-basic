package main

import "fmt"

func main() {
	// 参照型はslもsl2も同じアドレスにあるので、両方更新されてしまう
	// sl := []int{100, 200}
	// sl2 := sl

	// sl2[0] = 1000
	// fmt.Println(sl)

	sl := []int{1, 2, 3, 4, 5}
	sl2 := make([]int, 5, 10)
	fmt.Println(sl2)
	n := copy(sl2, sl)
	fmt.Println(n, sl2)

}
