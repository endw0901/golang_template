package main

import "fmt"

/*
▼main
・GOは自動的にmainをcallするため、明記不要

▼init
・mainと同様、引数無し・返り値無し
・mainの前に自動的にcallされる

▼initの利用用途
・グローバル変数の初期化
・functionの外側で宣言は可能だが、forループでの初期値設定はできない => init内で行う
*/

// mainの前に初期化したい
var numbers [10]int

func init() {
	fmt.Println("This is init#1")
}

func init() {
	fmt.Println("This is init#2")
	for i := 0; i < len(numbers); i++ {
		numbers[i] = i * 2
	}
}

func main() {

	// 1.initが自動でcallされる
	fmt.Println("This is main")

	// 2.initをcallすると未定義エラーとなる
	// init()

	// 3.initを重複して2つ定義すると、定義した順でcallされる
	/*
		This is init#1
		This is init#2
		This is main
	*/

	// 4.mainの前に初期化・初期値設定したglobal変数
	fmt.Printf("%#v\n", numbers) // [10]int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}
}
