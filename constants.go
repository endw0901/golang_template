package main

import (
	"fmt"
)

func main() {
	const days int = 7

	var i int
	fmt.Println(i)

	const pi float64 = 3.14
	const secondsInHour = 3600

	duration := 234 // in hours
	fmt.Printf("Duration in seconds: %v\n", duration*secondsInHour)

	// 1.変数ではruntimeエラー(コンパイル時点ではエラーにならず、実行時にエラー発生)
	// x, y := 5, 0
	// fmt.Println(x / y) // runtime error -> "panic: runtime error: integer divide by zero"

	// 2.constantsではコンパイル時点で早期エラー発見
	const a = 5
	const b = 0

	// fmt.Println(a / b) →コンパイルエラー

	const n, m int = 4, 5
	const n1, m1 = 6, 7

	// 3.複数のconst
	const (
		min1 = -500
		min2 = -300
		min3 = 100
	)

	fmt.Println(min1, min2, min3)

	// 4.省略すると前のconstのタイプと値を引き継ぐ
	const (
		min4 = -500
		min5
		min6
	)

	fmt.Println(min4, min5, min6)

	// 5.ルール(1)変更不可
	const temp int = 100
	// temp = 50  // コンパイルエラー

	// 6.ルール(2)runtimeに初期化は不可
	// const power = math.Pow(2, 3) // ※math.Powはbuilt in関数ではないため、runtimeに使われる

	// 7.ルール(3)変数で初期化は不可
	// t := 5
	// const tc = t // 変数からconstの設定がruntimeとなるためエラー

	// 8.ルール(4)len関数で初期化は可能 ※lenはbuilt in関数のため、コンパイル時に使える
	const l1 = len("hello")
	fmt.Println(l1)
}
