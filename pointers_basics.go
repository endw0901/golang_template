package main

import (
	"fmt"
)

/*
▼メモリについて
・RAM(computer memory)
・アドレスはhexadecimal(16進数) : 0xc00xxx
・変数はメモリアドレスの別名で、アドレスを保持する

▼ポインタについて
・&x： xのポインタ
・%p：fmt.Printfでポインタ表示
*/

func main() {

	// 1.変数のポインター(address)を表示する
	name := "Andrei"
	fmt.Println(&name) // 0xc00002e1f0

	// 2.ポインタの変数格納、Printf表示
	var x int = 2
	ptr := &x // *int (pointer to int)
	fmt.Println(&x)
	fmt.Printf("ptr is of type of %T with a value of %v and address %p\n", ptr, ptr, &ptr)
	/* ptr is of type of *int with a value of 0xc0000100d0 and address 0xc000006030 */

	fmt.Printf("address of x is %p\n", &x)

	// 3.ポインタタイプの変数
	var ptr1 *float64 // zero value is nil
	_ = ptr1

	// 4.newで変数作成
	p := new(int)
	x = 100
	p = &x

	fmt.Printf("p is a type of %T with a value of %v\n", p, p) // p is a type of *int with a value of 0xc0000100d0
	fmt.Printf("address of p is %p\n", &x)

	// 5.ポインタタイプ変数による参照: *p *x
	*p = 90            // x=90と同じことになる ※*pはxのアドレスを保持するので、xが指すアドレス先に格納すると等しい
	fmt.Println(x, *p) // 90 90

	fmt.Println("*p==x:", *p == x) // true

	/*
	   ▼違いがある点を理解する
	   *float64 : tyep description
	   *p : ポインター変数

	   &value   => pointer
	   *pointer => value
	*/

	*p = 10     // x = 10と同じ
	*p = *p / 2 // x / 2と同じ

	fmt.Println(x) // 5

}
