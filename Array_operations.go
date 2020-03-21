package main

import (
	"fmt"
	"strings"
)

/*
memo
*/

func main() {

	// 1.配列要素の更新
	numbers := [3]int{10, 20, 30}
	fmt.Printf("%#v\n", numbers)

	numbers[0] = 7
	fmt.Printf("%#v\n", numbers)

	// numers[5] = 100 // エラー:配列未定義

	// 2.iterationで配列にアクセス
	for i, v := range numbers {
		fmt.Println("index:", i, "value:", v)
	}

	fmt.Println(strings.Repeat("#", 20)) // 区切り文字

	for i := 0; i < len(numbers); i++ {
		fmt.Println("index:", i, "value:", numbers[i])
	}

	fmt.Println(strings.Repeat("#", 20)) // 区切り文字

	// 3.二次元配列
	balances := [2][3]int{
		{5, 6, 7},
		[3]int{8, 9, 10}, //コンマは必須
	}
	fmt.Println(balances)

	fmt.Println(strings.Repeat("#", 20)) // 区切り文字
	// 4.配列のコピー
	m := [3]int{1, 2, 3}
	n := m
	fmt.Println("n is equal to m:", n == m)

	m[0] = -1
	fmt.Println("n is equal to m:", n == m)

	/**
		arrayでは片方のみ変更される
	　　sliceでは両方変更される
		**/
}
