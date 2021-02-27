package main

import "fmt"

func main() {
	sl := []int{100, 200}
	fmt.Println(sl)

	sl = append(sl, 300)
	fmt.Println(sl)

	sl = append(sl, 400, 500, 600)
	fmt.Println(sl)

	sl2 := make([]int, 5)
	fmt.Println(sl2)

	fmt.Println(len(sl2))

	fmt.Println(cap(sl2))

	sl3 := make([]int, 5, 10)

	fmt.Println(len(sl3))
	//len = 要素数
	//cap = メモリの要領

	// 	【キャパシティ（要領）について】
	// 要領以上の要素が追加されるとメモリの消費が倍になってしまいます。
	// 過剰にメモリを確保してしまうと実行速度が落ちたりする。

	// 良質なパフォーマンスを実現するには、要領の管理も気にする。
	fmt.Println(cap(sl3))

}
