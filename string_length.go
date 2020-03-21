package main

import (
	"fmt"
	"unicode/utf8"
)

/*
unicode = 1～4bytes
ASCII = 1 byte
そのほかのunicode point はもっとbyteを使う

*/

func main() {
	// 1.ASCII code:1byte
	s1 := "Golang"
	fmt.Println(len(s1)) // 6文字 -> 6byte

	// 2.non-ASCII codeのbyte
	name := "田中"
	fmt.Println(len(name)) // 6 = 3byte*2

	// 3. rune count
	n := utf8.RuneCountInString(name) // 2 rune
	m := utf8.RuneCountInString("ゐ")  // 1 rune
	fmt.Println(n, m)
}
