package main

import "fmt"

/*
memo
*/

func main() {

	// 1.arrayの宣言方法
	var numbers [4]int           // 0で初期化される
	fmt.Printf("%v\n", numbers)  // [0 0 0 0]
	fmt.Printf("%#v\n", numbers) // [4]int{0, 0, 0, 0}

	// 2.値指定でarray宣言
	var a1 = [4]float64{}
	fmt.Printf("%#v\n", a1)

	var a2 = [3]int{-10, 1, 100}
	fmt.Printf("%#v\n", a2)

	a3 := [4]string{"Dan", "Diana", "Paul", "John"}
	fmt.Printf("%#v\n", a3)

	// 3.一部のみ初期化
	a4 := [4]string{"x", "y"}
	fmt.Printf("%#v\n", a4) // [4]string{"x", "y", "", ""} 残りは初期化される

	// 4. arrayのlengthを自動で判定：ellipsis operator [...]
	a5 := [...]int{1, 2, 3, 4, 5, 6, -10}
	fmt.Printf("%#v\n", a5)
	fmt.Printf("the length of a5 is %d\n", len(a5))

	a6 := [...]int{1,
		2,
		3,
		4,
		5, // 最後のカンマ必須
	}

	fmt.Printf("%#v\n", a6)
}
