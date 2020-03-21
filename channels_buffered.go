package main

import (
	"fmt"
	"time"
)

/*
channelは2種類存在する
(1)unbuffered channel
  => synchronous channelとも呼ばれる

(2)buffered channel
　=> send / receiveは分離される
     buffered size固定のため、上限が分かる


▼どのように使い分けるのか？buffered sizeはどう決めるのか？
buffered channelは、1つ前の工程が終わるまで待つ挙動
もし最初の工程の進捗が速い場合、待ちが増えることになるためbufferedのメリットがない

*/

func main() {

	// 1.unbuffered channel
	// c1 := make(chan int)

	// 2.buffered channel ※3つまでdataをbufferに持てる =>いっぱいになると空きがでるまでsendがblockされる
	c := make(chan int, 3)

	// 3.anonymous関数でchannelに値を送る
	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Printf("func goroutine #%d starts sending data into the channel\n", i)
			c <- i                                                                    // mainがwake upするまでfuncをblockする
			fmt.Printf("func goroutine #%d after sending data into the channel\n", i) // main goroutineがdataをreceiveするまでblockされる
		}
		close(c) // channelを閉じる　※全てのdataがchannelに送られたことを関数に伝えたいときにのみ必要
	}(c) // buffered channelを送る

	// 4.複数のgoroutineが起動する時間を作るため、mainはsleepする
	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	// 5.channelのデータをループ処理で出力する
	for v := range c { // v:= <-cのrangeループ版
		fmt.Println("main goroutine received value from channel:", v)
	}

	// 6.channelが空になった後で出力するとzero valueが出る=> intなので0
	fmt.Println(<-c) // 0

	// 7.closed channelにsendするとpanic errorとなる
	// c <- 10 //panic: send on closed channel

	/*
	   $ go run channels_unbuffered.go
	   main goroutine sleeps for 2 seconds
	   func goroutine starts sending data into the channel
	   main goroutine starts receiving data
	   main goroutine starts receiving data: 10
	   func goroutine after sending data into the channel


	   $ go run channels_buffered.go
	   main goroutine sleeps for 2 seconds
	   func goroutine #1 starts sending data into the channel
	   func goroutine #1 after sending data into the channel
	   func goroutine #2 starts sending data into the channel
	   func goroutine #2 after sending data into the channel
	   func goroutine #3 starts sending data into the channel
	   func goroutine #3 after sending data into the channel
	   func goroutine #4 starts sending data into the channel // mainが受け取る前にchannelにデータを送る
	   main goroutine received value from channel: 1  // FIFOで出力
	   main goroutine received value from channel: 2
	   main goroutine received value from channel: 3
	   main goroutine received value from channel: 4
	   func goroutine #4 after sending data into the channel // channel buffer=3なので、3data送ったらblockされ、空になったら4つめが送られる
	   func goroutine #5 starts sending data into the channel
	   func goroutine #5 after sending data into the channel
	   main goroutine received value from channel: 5
	   0 // channelが空になった後
	*/
}
