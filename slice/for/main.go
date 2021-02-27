package main

import "fmt"

// func main() {
// sl := []string{"A", "B", "C"}
// fmt.Println(sl)

// // for i, v := range sl {
// // 	fmt.Println(i, v)
// // }

// for i := 0; i < len(sl); i++ {
// 	fmt.Println(sl[i])
// }
// }

// 可変長引数 = 無限に引数を渡すことができる

func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}
func main() {
	fmt.Println(Sum(1, 2, 3))
	fmt.Println(Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
