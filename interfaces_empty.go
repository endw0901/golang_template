package main

import "fmt"

/*
・empty interfaceはGOのキーコンセプト
・どのタイプでも表すことができる

・空インターフェイスを使っているPrintln
func Println(a ...interface{})(n int, err error)
　=> どのタイプの引数を何個でも受け取れる

・Empty Interfaceは、タイプが不定のときに使われる
  関数の引数としてEmpty interfaceを引き渡すと、どのタイプでも受け取れる

・メンテが大変というデメリットあり
(タイプセーフコードの方が安全。タイプを明示したほうが手堅い)
　本当に必要な時だけ使用すること

*/

// 空のインターフェイス：どのタイプも実装できる
type emptyInterface interface {
}

// 空インターフェイスを1つだけ持つstruct
type person struct {
	info interface{} // empty interface
}

func main() {
	// 0.Empty Interfaceの宣言
	var empty interface{}

	// 1.数値を保持する空インターフェイス
	empty = 5
	fmt.Println(empty) // 5

	// 2.文字列も保持する空インターフェイス
	empty = "Go"
	fmt.Println(empty) // Go

	// 3.Sliceも保持する空インターフェイス
	empty = []int{4, 5, 6}
	fmt.Println(empty) // [4 5 6]

	// 4.ダイナミックタイプは長さ不定：assertionが必要
	// fmt.Println(len(empty)) // invalid argument empty (type interface {}) for len

	// 5.assertionを使ってダイナミックタイプの値を取得
	fmt.Println(len(empty.([]int))) // 3 ※sliceのlength

	// 6.空インターフェイスを保持するためどのタイプも保持できるstruct
	you := person{} // 空インターフェイスinfoを保持するperson structには、どのタイプも保持できる

	you.info = "Your name"
	fmt.Println(you.info) // Your name

	you.info = 40
	fmt.Println(you.info) // 40

	you.info = []float64{5.6, 5., 7.8}
	fmt.Println(you.info) // [5.6 5 7.8]

	// 7.
}
