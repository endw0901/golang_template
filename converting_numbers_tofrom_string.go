package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 1. 数字(int)を文字に変換：ASCII
	s := string(99) // ASCII code 99 = c
	fmt.Println(s)

	// 2. 数字(int)を文字に変換：fmt.Sprintf
	// s1 := string(44.2)
	var myStr = fmt.Sprintf("%f", 44.2)
	fmt.Println(myStr)

	var myStr1 = fmt.Sprintf("%d", 34234)
	fmt.Println(myStr1)

	// 3. 数字を文字に変換できない：unicode character : UTF8
	fmt.Println(string(34234)) // 数字を文字に変換できない。UTF8に変換される

	// 4. 数字(Float64)を文字に変換：strconv.Parse・・
	s1 := "3.123"
	fmt.Printf("%T\n", s1)

	var f1, err = strconv.ParseFloat(s1, 64) // Float64
	_ = err
	fmt.Println(f1 * 3.4)

	// 5.文字を数字に変換：ASCII to int
	i, err := strconv.Atoi("-50")

	// 6.数字を文字に変換：int to ASCII
	s2 := strconv.Itoa(20)

	fmt.Printf("i type is %T, i value is %v\n", i, i)
	fmt.Printf("s2 type is %T, s2 value is %q\n", s2, s2) // %q:引用文字
}
