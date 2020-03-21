package main

import "fmt"

func main() {
	var a = 4
	var b = 5.2

	// 1. floatをint変数に格納
	a = int(b)
	fmt.Println(a, b)

	// var x int
	// x = "5"

	// 2.初期値(初期値設定不要。初期化無しの変数はない)
	var value int
	var price float64
	var name string
	var done bool
	fmt.Println(value, price, name, done)
}
