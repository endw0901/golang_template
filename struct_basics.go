package main

import (
	"fmt"
)

/*
Brand : string -> "FORD"
BuiltYear:int -> 2019
Price:float64 -> 20,000

structはblueprint / dataの構造体
OOP(object oriented programming)のようなclass architectureはない。

*/

func main() {
	// 0.bookを変数で表すのは不便
	title1, author1, year1 := "The Devine Comedy", "Dante Aligheri", 1320
	title2, author2, year2 := "Macbeth", "William Shakespeare", 1606

	fmt.Println("Book1:", title1, author1, year1)
	fmt.Println("Book2:", title2, author2, year2)

	// 1.structタイプのbookを宣言
	type book struct {
		title  string
		author string
		year   int
	}

	// 2.同じタイプをまとめて宣言
	type book1 struct {
		title, author string
		year          int
	}

	// 3.book structを使う
	/*
		順番が重要になってしまう
		1.title, 2.author, 3:published year
	*/
	myBook := book{"The Dvine Comedy", "Dante Aligheri", 1320}
	fmt.Println(myBook) // {The Dvine Comedy Dante Aligheri 1320}

	// 4.順番関係なく名前で作成する
	bestBook := book{title: "Animal Farm", author: "George Orwell", year: 1945} // yearとauthorが逆だが問題なし
	_ = bestBook

	// 5.一部のみ指定したstruct -> 未宣言の項目は初期化される
	aBook := book{title: "Just a random book"}
	fmt.Printf("%#v\n", aBook) //main.book{title:"Just a random book", author:"", year:0} ※未宣言のauthor・yearには初期値が入る
}
