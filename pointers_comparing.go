package main

import "fmt"

/*
   *float64 : tyep description
   *p : ポインター変数

   &value   => pointer
   *pointer => value
*/

func main() {
	// 1.pointer格納した変数を別の変数に格納した場合
	a := 5.5
	p1 := &a // p1はaのアドレス(pointer)を保持
	pp1 := &p1

	fmt.Printf("Value of a: %v, address of a: %v\n", a, &a)
	fmt.Printf("Value of p1: %v, address of p1: %v\n", p1, &p1)
	fmt.Printf("Value of pp1: %v, address of pp1: %v\n", pp1, &pp1)

	/*
		aのアドレスはp1のvalue、p1のアドレスはpp1のvalueとなっている
			Value of a: 5.5, address of a: 0xc0000100b0
			Value of p1: 0xc0000100b0, address of p1: 0xc000006028
			Value of pp1: 0xc000006028, address of pp1: 0xc000006030
	*/

	// 2.変数に格納したアドレスが示す値を表示する (*x を使う)
	fmt.Printf("*p1 is %v\n", *p1)
	fmt.Printf("*pp1 is %v\n", *pp1)
	/*
	*p1 is 5.5 　※p1のアドレス(aのアドレス)が示すvalue(=aのvalue)
	*pp1 is 0xc0000100b0 ※pp1のアドレス(p1のアドレス)が示すvalue(=p1のvalue)
	 */

	// 3.**pp1 : 変数に格納したアドレスが示す値をアドレスと扱い、そのアドレスが示す値を表示する
	fmt.Printf("**pp1 is %v\n", **pp1) // **pp1 is 5.5
	/*
	 *pp1 ： pp1のアドレス(p1のアドレス)が示すp1のvalueのアドレス(=aのアドレス)が示すvalue(=aのvalue)を表示する
	 */

	// 4. **pp1をつかって、pp1->p1->aとたどってaの値を変更する
	**pp1++
	fmt.Printf("a is %v\n", a) // a is 6.5

	// 5.初期化無しのpointer変数 (nil)
	var p2 *int
	fmt.Printf("%#v\n", p2) // (*int)(nil)

	// 6.pointerの比較
	y := 5
	p2 = &y

	z := 5
	p3 := &z

	fmt.Println(p2 == p3) // false

	p4 := &y
	fmt.Println(p2 == p4) // true
}
