package main

import (
	"fmt"
	"time"
)

/*
・GOにclassはないが、methodがある
・type + extra behavior
*/

// method用のslice =>引数となる
type names []string // slice

// method
func (n names) print() {
	for i, name := range n {
		fmt.Println(i, name)
	}
}

func main() {

	// 1.既存のnamed methodの例：time. day.
	const day = 24 * time.Hour
	fmt.Printf("%T\n", day) // time.Duration

	seconds := day.Seconds()
	fmt.Printf("%T\n", seconds)                   // float64
	fmt.Printf("Seconds in a day: %v\n", seconds) // Seconds in a day: 86400

	// 2.独自のnamed method
	friends := names{"Dan", "Marry"}
	friends.print() // methodの場合、引数.method名の順に記載
	/*
		0 Dan
		1 Marry
	*/

	// 3.methodの別のcall方法
	names.print(friends)

	/*
		▼methodとfunctionの違い
		method は typeに属する
		functionはpackageに属する
	*/

	// 4.基底タイプが同じタイプを変換して特定のmethodにアクセス
	var n int64 = 93422433
	fmt.Println(n)                // 93422433
	fmt.Println(time.Duration(n)) //int64をコンバート => 93.422433ms
}
