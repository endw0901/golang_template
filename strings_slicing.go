package main

import (
	"fmt"
)

/*


 */

func main() {
	// 1. ASCIIのslicing -> byteを返す
	s1 := "I love Golang"
	fmt.Println(s1[2:5]) // lov 3文字目～5文字目を返す(2-4)

	// 2. non-ASCIIのslicing -> byteを返す
	s2 := "内閣総理大臣官邸"
	fmt.Println(s2[0:2]) // �� ※2bytes ※この2byteのunicode表示がこの文字化けとなる

	// 3. slice of runeを返す（byteではなく)
	rs := []rune(s2)
	fmt.Printf("%T\n", rs) // rune is alias of int32

	fmt.Printf(string(rs[0:3])) //内閣総
}
