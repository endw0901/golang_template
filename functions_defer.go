package main

import (
	"fmt"
	"log"
	"os"
)

/*
defer
上位の(surrounding)関数が処理を終えるまで遅らせる(deferする)

・最後にcallされることを保証するために使う　例：fileのclose
  mainの終了後にcloseされる。
*/

func foo() {
	fmt.Println("This is foo()")
}

func bar() {
	fmt.Println("This is bar()")
}

func foobar() {
	fmt.Println("This is foobar()")
}

func main() {

	// 0.通常の関数call
	foo()
	bar()
	/*
		This is foo()
		This is bar()
	*/

	// 1.defer
	defer foo() // 上位関数 = main関数が終わるまで実行を遅らせる(deferする)
	bar()
	fmt.Println("Just a string after defering foo() and calling bar()")
	/*
	   This is bar()
	   Just a string after defering foo() and calling bar()
	   This is foo()
	*/

	// 2.複数のdefer -> 順番が逆にcallされる
	defer foobar()

	/*
		This is bar()
		Just a string after defering foo() and calling bar()
		This is foobar()
		This is foo()
	*/

	// 3.deferの使用目的を示す例：file.Close()
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// code that works (read/write) with the file
}
