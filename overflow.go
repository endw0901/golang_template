package main

import (
	"fmt"
	"math"
)

func main() {

	// 1.overflow -> 最小値にリセットする(0)
	var x uint8 = 255
	x++
	fmt.Println(x)

	// 2.コンパイル時にエラー

	// a := int8(255 + 1)

	// 3.overflow -> 最小値にリセットする(-128) -> -1すると127に戻る
	var b int8 = 127

	fmt.Printf("%d\n", b+1)

	b = -128
	b--
	fmt.Println(b)

	// 4.overflow -> 無限
	f := float32(math.MaxFloat32)
	fmt.Println(f)

	f = f * 1.2
	fmt.Println(f) // +Inf = infinite(無限)

	// 5.constantとoverflow
	// const i int8 = 300 // コンパイルエラー
}
