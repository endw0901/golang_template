package main

import (
	"fmt"
	"strings"
)

/*
 */

// variadic function
func f1(a ...int) {
	fmt.Printf("%T\n", a)
	fmt.Printf("%#v\n", a)
}

func f2(a ...int) {
	a[0] = 50 // 最初の要素を変更
}

func SumAndProduct(a ...float64) (float64, float64) {
	// sliceを受けるvariadic関数 -> float64を2つ返す
	sum := 0.
	product := 1.
	for _, v := range a { // sliceのiteration
		sum += v
		product *= v
	}
	return sum, product
}

// non-variadicとvariadic引数の混在
func personInfomation(age int, names ...string) string {
	fullName := strings.Join(names, " ") // 空白を区切り文字としてコンカチ

	returnString := fmt.Sprintf("Age: %d, fullName: %s", age, fullName) // 返り値を編集
	return returnString                                                 // 指定した変数を返り値に指定
}

func main() {
	// 1. variadic function 引数の数が不定
	f1(1, 2, 3, 4) // []int{1, 2, 3, 4}

	// 2.variadic関数に何も渡さないとnilで初期化される
	f1() // []int(nil)

	// 3.variadic関数の例：append ※引数にいくつでもつけられる
	nums := []int{1, 2}
	nums = append(nums, 3, 4)

	// 4.variadic関数にslice + variadic operator(...)
	f1(nums...) // []int{1, 2, 3, 4}

	// 5.variadic関数にsliceを引き渡し、sliceの要素の一つを更新する
	f2(nums...)
	fmt.Println("Nums:", nums) // Nums: [50 2 3 4]

	// 6.variadic関数に不定数の引数を引き渡す例(関数内でsliceのiterationで演算)
	s, p := SumAndProduct(2.0, 5., 10.) // 10.は、10のfloat valueという指定
	fmt.Println(s, p)                   // 17(2+5+10) 100(2*5*10)

	s, p = SumAndProduct(2.0, 5., 10., 5.6, 5.6, 87.3)
	fmt.Println(s, p) // 115.5 273772.8

	// 7.non-variadic + variadicの順で指定すると引数を混ぜることができる
	info := personInfomation(30, "Wolfgang", "Amadeus", "Mozart")
	fmt.Println(info) // Age: 30, fullName: Wolfgang Amadeus Mozart
}
