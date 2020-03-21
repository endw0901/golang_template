package main

import "fmt"

/*
・arrayは要素数固定、slicesは変動
・arrayは未指定で0初期化、slicesはnil
・arrayはコンパイル時に定義され、slicesはruntime(実行時)の定義
*/

func main() {

	// 1.宣言
	var cities []string // 未指定はnilで初期化される
	fmt.Println("cities is equal to nil: ", cities == nil)
	fmt.Printf("cities %#v\n", cities)
	fmt.Println(len(cities)) // len = 0

	// cities[0] = "London" //エラー index out of range

	numbers := []int{2, 3, 4, 5}
	fmt.Println(numbers)

	// 2.makeによる宣言 (make:built in function)
	nums := make([]int, 2)
	fmt.Printf("%#v\n", nums) // []int{0, 0}

	// 3.sliceの特定の要素に値設定
	type names []string
	friends := names{"Dan", "Maria"}
	fmt.Println(friends)

	myFriend := friends[0]
	fmt.Println("My best friend is ", myFriend)

	friends[0] = "Gabriel"
	fmt.Println("My best friend is ", friends[0])

	// 4. iteration
	for index, value := range numbers {
		fmt.Printf("index: %v, value: %v\n", index, value)
	}

	// 5. sliceのコピー
	var n []int
	n = numbers
	fmt.Println(n)

}
