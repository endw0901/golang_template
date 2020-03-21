package main

import (
	"fmt"
	"time"
)

/*
channelは2種類存在する
(1)buffered channel
(2)unbuffered channel
*/

func main() {

	// 1.unbuffered channel
	c1 := make(chan int)

	// 2.buffered channel
	c2 := make(chan int, 3)
	_ = c2

	// 3.anonymous関数でchannelに値を送る
	go func(c chan int) {
		fmt.Println("func goroutine starts sending data into the channel")
		c <- 10                                                           // mainがwake upするまでfuncをblockする
		fmt.Println("func goroutine after sending data into the channel") // main goroutineがdataをreceiveするまでblockされる
	}(c1)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	fmt.Println("main goroutine starts receiving data")
	d := <-c1
	fmt.Println("main goroutine starts receiving data:", d)

	time.Sleep(time.Second)

	/*
		main goroutine sleeps for 2 seconds
		func goroutine starts sending data into the channel
		main goroutine starts receiving data
		main goroutine starts receiving data: 10
		func goroutine after sending data into the channel
		※main goroutineがdataをreceiveするまでblockされる
	*/
}
