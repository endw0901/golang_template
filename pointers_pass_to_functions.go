package main

import "fmt"

/*
   *float64 : tyep description
   *p : ポインター変数

   &value   => pointer
   *pointer => value

   **pointer => value(pointer)のvalue
*/

// 入力ポインターの示す値を変更し、別のポインターを返す変数
func change(a *int) *float64 {
	*a = 100

	b := 5.5
	return &b // ほかの言語では問題がある場合があるが、GOではOK
}

func changeVar(a int) {
	a = 66
}

func main() {
	x := 8
	p := &x

	// 1.関数でポインターの変更、ポインターのリターン ※間接的な引き渡しで変更が可能となるメリット
	fmt.Println("value of x before calling change():", x) // 8
	change(p)
	fmt.Println("value of x after calling change():", x) // 100

	// 2.xの変数自体を引き渡して関数内で編集しても変更されない
	fmt.Println("value of x before calling changeVar():", x) // 100
	changeVar(x)
	fmt.Println("value of x after calling changeVar():", x) // 100
}
