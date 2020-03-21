package main

import "fmt"

/*
・map同士の比較は不可。nilと比較は可能

map[keyタイプ]valueタイプ{"要素","要素"}
*/

func main() {
	// 1.map同士は比較できない
	a := map[string]string{"A": "X"}
	b := map[string]string{"B": "Y"}
	// fmt.Println(a == b) // invalid operation: a == b (map can only be compared to nil)

	// 2.Sprintf
	s1 := fmt.Sprintf("%s", a)
	s2 := fmt.Sprintf("%s", b)

	fmt.Println(s1, s2) // map[A:X] map[B:Y]

	if s1 == s2 {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are not equal")
	}
}
