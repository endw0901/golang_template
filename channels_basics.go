package main

import (
	"fmt"
)

/*
・channel = 2つのgoroutine間のコネクション
・channelでやりとりするデータは同じタイプのみ
・channelは双方向の連絡オブジェクト (send <=> receive)
・やり取りはchannel operator( <- )で行われる

*/

func main() {
	// 1.channel宣言
	var ch chan int // channel間のデータやり取りはint
	fmt.Println(ch) // 初期値はnil(zero value) <nil>

	// 2.channelの初期化 => ポインタが格納される
	ch = make(chan int)
	fmt.Println(ch) // 0xc000040120

	// 3.channelの宣言＆初期化を同時に
	c := make(chan int)

	// 4.channel operator:(1)send operation
	c <- 10 // channelに値(10)を送る

	// 5.channel operator:(2)RECEIVE
	num := <-c // channelから値を受け取る

	// 6.channel operatorで受け取った値をそのまま出力
	fmt.Println(<-c)
	_ = num

	// 7.channelをclose => その後のsendはpanic erro => channelから値がなくなるまでreceiveは可能 => その後のreceiveは全てnil(zero value)
	close(c)

	// 8.この時点で実行してもfatal errorとなる(main goroutineしかないので通信できない：deadlock)
	/*
		fatal error: all goroutines are asleep - deadlock!

		goroutine 1 [chan send]:
		main.main()
		        C:/Users/masaya/go/src/master_go_programming/channels_basics.go:27 +0x146
		exit status 2
	*/
}
