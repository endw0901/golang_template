package main

import "fmt"

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

}
