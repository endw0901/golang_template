package main

import (
	"fmt"
)

/*
memo
*/

func main() {

	type book struct {
		title  string
		author string
		year   int
	}

	lastBook := book{title: "Anna Karenina"}

	// 1.structの指定した項目の値を取得
	fmt.Println(lastBook.title) // Anna Karenina

	// 2.structの存在しない項目を指定 -> エラー
	// fmt.Println(lastBook.pages) // エラー:lastBook.pages undefined (type book has no field or method pages)

	fmt.Printf("%#v\n", lastBook) // main.book{title:"Anna Karenina", author:"", year:0}

	// 3.update
	lastBook.author = "Leo Tolstoy"
	lastBook.year = 1878
	fmt.Printf("%+v\n", lastBook) // {title:Anna Karenina author:Leo Tolstoy year:1878}

	// 4.structの比較
	aBook := book{title: "Anna Karenina", author: "", year: 0}
	fmt.Println(aBook == lastBook) // false

	aBook = book{title: "Anna Karenina", author: "Leo Tolstoy", year: 1878}
	fmt.Println(aBook == lastBook) // true

	// 5.structのコピー:ポインタだけでなく、新しいメモリアドレスを作る
	myBook := aBook
	myBook.year = 2020         // aBookは変更されない
	fmt.Println(myBook, aBook) // {Anna Karenina Leo Tolstoy 2020} {Anna Karenina Leo Tolstoy 1878}

	// 6.struct type aliasを使わずに宣言
	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Diana",
		lastName:  "Muller",
		age:       30, // 最後のカンマ必須
	}

	fmt.Printf("%#v\n", diana)                 // struct { firstName string; lastName string; age int }{firstName:"Diana", lastName:"Muller", age:30}
	fmt.Printf("Diana's Age: %d\n", diana.age) // Diana's Age: 30

	// 7.field name無しに宣言(anonymous struct)
	type Book struct {
		string
		float64
		bool
	}

	b1 := Book{"1984 by George Orwell", 10.2, false}
	fmt.Printf("%#v\n", b1) // main.Book{string:"1984 by George Orwell", float64:10.2, bool:false}

	// 8.anonymous structの項目値取得(typeを指定(field nameがないので))
	fmt.Println(b1.string) // 1984 by George Orwell

	// 9.一部にfield nameがあって、一部に無いstruct
	type Employee struct {
		name   string
		salary int
		bool   // 初期化：false
	}

	e := Employee{"John", 40000, false}
	fmt.Printf("%#v\n", e) // main.Employee{name:"John", salary:40000, bool:false}

	e.bool = true // 更新
	fmt.Printf("%#v\n", e)
}
