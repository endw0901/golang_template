package main

import "fmt"

func main() {
	/*
	   unitは、unsigned int→正の整数のこと
	*/
	// 1.INT (int8=保存に8bit使う)
	var i1 int8 = -128 // -129でoverflow
	fmt.Printf("%T\n", i1)

	var i2 uint16 = 65535 // max ※越えるとoverflowエラー
	fmt.Printf("%T\n", i2)

	// 2.FLOAT
	var f1, f2, f3 float64 = 1.1, -.2, 5. // 1.1, -0.2, 5.0
	fmt.Printf("%T %T %T\n", f1, f2, f3)

	// 3.RUNE
	// ※文字列 -> byte / rune

	var r rune = 'f' //102 = decimal ASCIIコードの'f'
	fmt.Printf("%T\n", r)
	fmt.Println(r)

	fmt.Printf("%x\n", r) //102 = hexadecimal ASCIIコードの'f'

	// 4.BOOL
	var b bool = true
	fmt.Printf("%T\n", b)
	fmt.Println(b)

	// 5.STRING
	var s string = "Hello GO!"
	fmt.Printf("%T\n", s)
	fmt.Println(s)

	// 6.ARRAY
	/*
	 ARRAY = 数が決まっている
	 SLICE = 数が変動する
	 ※どちらも配列の中の全要素は同じタイプ
	*/

	var numbers = [4]int{4, 5, -9, 100}
	fmt.Printf("%T\n", numbers)

	// 7.SLICE
	var cities = []string{"London", "Tokyo", "New York"}
	fmt.Printf("%T\n", cities)

	// 8.MAP
	/*
		Pythonのdictionaryに似ている
		・順不同
		・全要素が同じタイプ
		・他のタイプのキーを持つ
	*/

	balances := map[string]float64{
		"USD": 2332.2,
		"EUR": 511.11,
	}
	fmt.Printf("%T\n", balances)

	// 9.STRUCT
	/*
		・順番
		・オブジェクト指向
	*/

	type Person struct {
		name string
		age  int
	}

	var you Person
	fmt.Printf("%T\n", you)

	// 10.POINTER
	/*
		・メモリーアドレスを保存
		・初期化無し = nil (ニル)
	*/

	var x int = 2
	ptr := &x // xのアドレス
	fmt.Printf("ptr is of type %T wiht the value of %v\n", ptr, ptr)

	// 11.FUNCTION TYPE
	fmt.Printf("%T\n", f)

	// 残り・・・Interface, Channel → 別にまとめる
}

func f() {

}
