package main

import "fmt"

/*
memo
*/

func main() {

	// 1.初期化無し：nil
	var n []int
	fmt.Println(n == nil) // true

	// 2.初期化するが空:nilではない
	m := []int{}
	fmt.Println(m == nil) // false

	// 3.slicesの比較 は==ではできない
	a, b := []int{1, 2, 3}, []int{1, 2, 4}
	// fmt.Println(a == b) // invalid operation: a == b (slice can only be compared to nil)

	// 4.slicesの比較→iterationで行う
	var eq bool = true

	for i, valueA := range a {
		if valueA != b[i] { // a[i]とb[i]を比較
			eq = false
			break
		}
	}

	if eq {
		fmt.Println("a and b slices are equal")
	} else {
		fmt.Println("a and b slices are not equal")
	}

	// 5.nilにすると比較ができなくなる
	c, d := []int{1, 2, 3}, []int{1, 2, 4}

	var eq2 bool = true

	c = nil // nilにすると比較ができないので、初期値のtrueのまま
	for i, valueA := range c {
		if valueA != d[i] {
			eq2 = false
			break
		}
	}

	if len(c) != len(d) { // nilにすると比較ができないので、lenの比較をすることで異なるということが分かる
		eq2 = false 
	}

	if eq2 {
		fmt.Println("c and d slices are equal")
	} else {
		fmt.Println("c and d slices are not equal")
	}
}
