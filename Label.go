package main

import "fmt"

/*
(1)break,continue,gotoで使う
(2)使われないlabelはエラー
(3)他の識別子と競合しない。Label内でユニークならok
(4)labelが宣言されたfunction内で有効
(5)主にネストループの外側のループを終了するために使う
*/

func main() {

	// 1.基本

	outer := 19 // Labelは他の識別子（同じ名前の変数)とは競合しない
	_ = outer

	people := [5]string{"Hellen", "Mark", "Brenda", "Antonio", "Michael"}
	friends := [2]string{"Mark", "Marry"}

outer: // ラベル

	for index, name := range people {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("found a friend %q at index %d \n", friend, index)
				break outer // ラベル無しだとfrenadのloopのみ抜ける
			}
		}

	}
	fmt.Println("Next instruction after the break")

}
