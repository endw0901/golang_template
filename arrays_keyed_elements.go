package main

import "fmt"

/*
・indexはオプション
・正の整数
・指定がなければ、最初は0、次はひとつ前のindex +1
*/

func main() {

	// 1.index
	grades := [3]int{
		1: 10,
		0: 5,
		2: 7,
	}
	fmt.Println(grades) // [5 10 7]

	// 2.一部のみindex指定
	accounts := [3]int{2: 50} // [0 0 50]
	fmt.Println(accounts)

	names := [...]string{
		5: "Dan",
	}
	fmt.Println(names, len(names)) // [     Dan] 6

	cities := [...]string{
		5:        "Paris",
		"London", // 指定しないとき、ひとつ前のindex + 1となる
		1:        "NYC",
	}
	fmt.Printf("%#v\n", cities) // [7]string{"", "NYC", "", "", "", "Paris", "London"}

	// 3.
	weekend := [7]bool{5: true, 6: true}
	fmt.Println(weekend) // [false false false false false true true]
}
