package main

import "fmt"

/*
memo
*/

func main() {

	// 1.append 要素追加
	numbers := []int{2, 3}

	numbers = append(numbers, 10)
	fmt.Println(numbers) // [2 3 10]

	numbers = append(numbers, 20, 30, 40)
	fmt.Println(numbers)

	// 2.他のsliceを追加
	n := []int{100, 200}
	numbers = append(numbers, n...) // n... elipsis operator
	fmt.Println(numbers)

	// 3.copy slice
	src := []int{10, 20, 30}
	dst := make([]int, len(src)) // 要素数が同じ
	nn := copy(dst, src)
	fmt.Println(src, dst, nn) // [10 20 30] [10 20 30] 3

	// 4.copy slice ※要素数が違う時
	src2 := []int{10, 20, 30}
	dst2 := make([]int, 2) // lenが異なるsliceにコピー
	nn2 := copy(dst2, src2)
	fmt.Println(src2, dst2, nn2) // [10 20 30] [10 20] 2
}
