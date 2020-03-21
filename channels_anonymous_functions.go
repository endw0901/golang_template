package main

import (
	"fmt"
	"strings"
)

/*
go f1 とmainの間でchannel通信するシンプルPGM
*/

// factorial of N関数(2～N(入力)まで掛ける) 2*3*4*5=120
func factorial(n int, c chan int) {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}

	// 計算後のfをchannelに送る
	c <- f
}

func main() {

	// 1.典型的なchannel宣言とcloseのベストプラクティス
	ch := make(chan int) // channel宣言＆初期化を同時に行う
	defer close(ch)      // mainの終了までchannelのcloseを遅らせる(deferする)

	// 2.goroutineを実行(channel通信付き)
	go factorial(5, ch)

	// 3.blocking call：mainはメッセージが関数から来るまで待ちたい
	f := <-ch      // blocking call : channelからメッセージが返るまでmainはスリープする
	fmt.Println(f) // 120 2*3*4*5(2～n(5)まで掛ける=factorial of n)

	// 4.多くの関数を動かして、本当にblocking callでmainが全て関数が終わるまで待つのかテストする
	for i := 1; i <= 20; i++ {
		go factorial(i, ch)
		f := <-ch
		fmt.Println(f)
	}

	fmt.Println(strings.Repeat("#", 10))

	// 5.anonymous関数
	for i := 5; i <= 15; i++ {
		go func(n int, c chan int) {
			f := 1
			for i := 2; i <= n; i++ {
				f *= i
			}

			// 計算後のfをchannelに送る
			c <- f
		}(i, ch) // i->n、ch(チャネル変数)-> c chan intへ送る。anonymous function ※go factorial(i,ch)のところと同じ

		fmt.Printf("Factorial of %d is %d\n", i, <-ch)
	}

	/*
		Factorial of 5 is 120
		Factorial of 6 is 720
		Factorial of 7 is 5040
		Factorial of 8 is 40320
		Factorial of 9 is 362880
		Factorial of 10 is 3628800
		Factorial of 11 is 39916800
		Factorial of 12 is 479001600
		Factorial of 13 is 6227020800
		Factorial of 14 is 87178291200
		Factorial of 15 is 1307674368000
	*/
}
