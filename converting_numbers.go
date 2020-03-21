package main

import "fmt"

func main() {

	/*
	   converting
	   他の言語ではcasting
	*/

	var x = 3
	var y = 3.1

	// x = x * y // タイプエラー

	// 1.convert
	x = x * int(y)
	fmt.Println(x)

	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", int(y))

	x = int(float64(x) * y)
	fmt.Println(x)

	y = float64(x) * y
	fmt.Println(y)

	// 2. int ≠ int64
	var a = 5
	fmt.Printf("%T\n", a)

	var b int64 = 2
	a = int(b)

}
