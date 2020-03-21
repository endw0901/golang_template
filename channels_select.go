package main

import (
	"fmt"
	"time"
)

/*
▼channels selectについて
caseのどれかが終わるまでblockする
複数の選択肢がある場合、ランダムにcaseを選ぶ

selectはchannelのみの機能

▼処理が居よう
・2秒ずつ待つ2つの関数でchannelに2つの文字列を送り、select caseで出力
 => 2秒で終了するため
 => selectでは2つの関数が同時に実行されているとわかる
*/

func main() {

	// 0.処理時間を計測する
	start := time.Now().UnixNano() / 1000000

	// 1.unbuffered channelを2つ宣言＆初期化する
	c1 := make(chan string)
	c2 := make(chan string)

	// 2.anonymous関数でchannelに値を送る
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "Hello!"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "Salut!"
	}()

	// 3.時間を作る
	/*
		これがないと、select caseはnon-blockingのためchannelが空のまま実行され、defaultを通り終了してしまう
	*/
	time.Sleep(time.Second * 2)

	// 4.select
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received:", msg1)
		case msg2 := <-c2:
			fmt.Println("Received:", msg2)
		default:
			fmt.Println("No activity")
		}
	}

	// 5.処理時間を計測する(終了)
	end := time.Now().UnixNano() / 1000000

	fmt.Println(start)
	fmt.Println(end)
	fmt.Println(end - start) // 20042

	/*
		goにはこれしかないので ナノ秒を100万で割ってミリ秒を作る
		1000ミリ秒 = 1秒
		2,042ミリ秒=2秒
	*/

	/*
		Received: Hello!
		Received: Salut!
		1584769291341
		1584769293343
		2002
	*/

	/*
		※select caseにdefaultがある場合(かつ、select前にsleepしていないとき)
		No activity
		No activity
		1584769504775
		1584769504776
		1
	*/
}
