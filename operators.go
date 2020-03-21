package main

import "fmt"

func main() {

	/*
	   1.数学Arithmatic operators
	   +
	   -
	   *
	   /
	   % modulus or mod → 余りを返す
	   ※modulusはintegerのみ(×float)

	   ※+はstringにも使える（コンカチ)
	*/

	// 1-1.四則演算
	a, b := 4, 2
	r := (a + b) / (a - b) * 2

	fmt.Println(r)

	// 1-2.余り
	r = 9 % a

	fmt.Println(r)

	// 2.Assignment Operators
	/*

				=
				+=
				-=
				*=
				/=
				%=

		++,--はstatement (operatorではない)

	*/

	a2, b2 := 2, 3
	a2 += b2 // 2 + 3
	fmt.Println(a2)

	b2 -= 2 // 3-2
	fmt.Println(b2)

	b2 *= 10
	fmt.Println(b2)

	b2 /= 5
	fmt.Println(b2)

	a2 %= 3
	fmt.Println(a2) // 2%3の余り2

	//3.++ -- statement (operatorではない)

	x := 1
	// x += 1 // 1+1=2
	x++ // 2+1=3
	fmt.Println(x)

	x-- // 3-2=2
	fmt.Println(x)

	// fmt.Println(5 + x--) // operatorではないのでエラー

	// 4.比較operators
	/*
		==
		!=  ※not equal
		<
		<=
		>
		>=
	*/

	a3, b3 := 5, 10
	fmt.Println(a3 == b3) // false
	fmt.Println(a3 != b3)
	fmt.Println(a3 > b3) // false

	// 5.論理operators
	/*
		&&   ※and
		||   ※or
		!
	*/

	a4, b4 := 5, 10
	fmt.Println(a4 > 1 && b4 == 10)   // true
	fmt.Println(a4 == 5 || b4 == 100) // true

	fmt.Println(!(a4 > 0))
	fmt.Println(!(a4 < 0))
	fmt.Println(!(a==1)||(b==100))
}
