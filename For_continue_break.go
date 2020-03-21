package main

import "fmt"

func main() {
	// 1.continueはloopを途中終了して、loopは継続する
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			// 奇数は抜けてcontinue
			continue
		}
		fmt.Println(i)
	}

	// 2.breakはloop自体を抜ける
	count := 0

	for i := 0; true; i++ {
		// trueは無限ループ
		if i%13 == 0 {
			// 13で割り切れる時
			fmt.Printf("%d is divisible by 13 \n", i)
			count++
		}
		if count == 10 {
			// 13で割り切れる数字を10個見つけたら無限forループを抜ける
			break
		}
	}
	fmt.Printf("after the loop")

}
