package main

import "fmt"

func main() {
	// 0.typed constant
	const a float64 = 5.1

	// 1.タイプ無し初期化
	const b = 6.7

	// 2.算出・コンカチで初期化
	const c float64 = a * b
	const str = "Hello " + "Go!"

	// 3. true/false 初期化
	const d = 5 > 10
	fmt.Println(d)

	// 4. int / floatの違いあり →タイプ無し初期化
	// const x int = 5
	// const y float64 = 2.2 * x // タイプの違いエラー

	const x = 5 // タイプ無しはintとなるが、floatに変換も可能となる(intで明示するとtypeエラー)
	const y = 2.2 * x

	fmt.Printf("%T\n", y)

	var i int = x     // x -> intに変わる
	var j float64 = x // x -> float64に変わる
	var p byte = x
	fmt.Println(i, j, p)

	const r = 5
	var rr = r
	fmt.Printf("%T", rr)

}
