package main

import "fmt"

/*
memo
他の言語同様、あまり推奨されない。注意して使う
*/

func main() {

	// 1.基本
	i := 0

loop: // Label
	if i < 5 {
		fmt.Println(i)
		i++
		goto loop
	}

	// 2.未使用の変数を越えてgotoはできない
	//	goto todo // xの宣言を飛ばしてLabelに飛ぶことはできない（エラー)
	//	x := 5
	// todo:
	//	fmt.Println("something here")

}
