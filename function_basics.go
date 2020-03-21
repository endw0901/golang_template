package main

import (
	"fmt"
	"math"
)

/*
Function
・推奨：camelCase
・同じpackage内ではfunction名はユニーク
・overloadingは不可
*/

// 1.
func f1() {
	fmt.Println("This is f1() function.")
}

// 2.
func f2(a int, b int) {
	fmt.Println("Sum:", a+b)
}

// 3.同じタイプの引数が続くとき -> 省略形
func f3(a, b, c int, d, e float64, s string) {
	fmt.Println(a, b, c, d, e, s)
}

// 4.返り値のタイプは省略可能（下記は明記した例)
func f4(a float64) float64 {
	return math.Pow(a, a)
}

// 6.
func f5(a, b int) (int, int) {
	return a + b, a * b
}

// func sum(a,b int)(s,a int, z float64, s, y string){ // これもok

// 7.
func sum(a, b int) (s int) {
	fmt.Println(s) // s-> 0
	s = a + b
	return // named returnにはreturnを使う
}

func main() {
	// 1.引数無し関数をcall
	/*
	   goはmainとinitのみ自動で実行する。
	   それ以外の関数はmain,init内からcallする
	*/
	f1()

	// 2.複数の引数を取る関数
	f2(3, 5)

	// 3.同じタイプの引数が続くとき -> 省略形
	f3(4, 5, 6, 2.1, 3.2, "golang")

	// 4.返り値のタイプは省略可能
	// 5.返り値を表示
	p := f4(2.4)
	fmt.Println(p)

	// 6.複数の返り値を持つ関数
	a, b := f5(8, 9)
	fmt.Println(a, b)

	// 7. 名前付き返り値：named return value
	mySum := sum(4, 8)
	fmt.Println(mySum) // 0, 12
}
