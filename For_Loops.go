package main

import "fmt"

func main() {
	// 1.基本
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 2.i++省略
	for i := 0; i < 10; {
		fmt.Println(i)
		i++
	}

	// 3.whileの動き(Goにwhile loopは無い)
	j := 10
	for j >= 0 {
		fmt.Println(j)
		j--
	}

	// 4.無限ループ　→ Ctrl + Cで強制終了
	// sum := 0
	// for {
	// 	sum++
	// }
	// fmt.Println(sum)

}
