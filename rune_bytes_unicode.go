package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/*
(1)integer
(2)byte = uint8 ※alias
(3)rune = int32 ※alias

・byte,runeはintegerとstringの区別のために使われる
・charはない。byte,runeが文字列
・ASCII encoded in UTF8 default

◆rune
・runeは''シングルクォート 'a' '\n'
・runeはUnicode Code Points ※rune literalを表す数値

128 code pointsでできるUnicodeのsubset = ASCII

・stringはslice of bytes

・rune 61 in hex = rune literal 'a'

*/

func main() {

	// 1.rune
	/**
	シングルクォーテーション:rune
	  → int32、
	  98は bのcode point
	**/
	var1, var2 := 'a', 'b'                          // rune (unicode code points)
	fmt.Printf("Type: %T, Value: %d\n", var1, var2) // Type: int32, Value: 98

	// 2.stringの中身
	str := "tara"
	fmt.Println(len(str)) // 4bytes

	fmt.Println("Byte(not rune at position 1:", str[1]) // 97

	// 3.stringの要素を取り出す
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i]) // %c = string or rune
	}

	// 4.stringをrune by runeでdecode（runeをstringとsizeで返す） (1)iteration (2)range
	fmt.Println("\n" + strings.Repeat("#", 20))

	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%c", r)
		i += size
	}

	fmt.Println("\n" + strings.Repeat("#", 20))
	for _, r := range str {
		fmt.Printf("%c", r)
	}

}
