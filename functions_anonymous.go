package main

import (
	"fmt"
)

/*
anonymous function
・nameを保持しない
・declared in line using function literal？
・form closuresに使われる？

・return a function inside another functionのユースケースで有用性が分かる
*/

func increment(x int) func() int { // intを返す別のfuncをreturnするfunc increment
	return func() int {
		x++
		return x
	}
}

func main() {
	// 1.function without a name
	func(msg string) { // anonymous function (has no name)
		fmt.Println(msg)
	}("I'm an anonymous function") // invoke an argument now

	// 2.functionをリターンするfunction
	// function can return another function
	a := increment(10)
	fmt.Printf("%T\n", a) // func() int -> aは「func() int」タイプなので、関数としてcallできる(a()) ※リターンはint

	a()              // 10+1=11
	fmt.Println(a()) //11+1=12
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
}
