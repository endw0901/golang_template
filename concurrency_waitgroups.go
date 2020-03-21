package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
・WaitGroupにより複数のgoroutineをシンクロさせる(f1が終わるまえにmainが終了する問題を解消)

▼使い方
(1)package "sync" をimportする
(2)関数：goroutineの引数にWaitGroupをポインタ引き渡しする
(3)関数：goroutine内で、goroutineが終了したことをWaitGroupに通知する処理を追加する : wg.Done()
(4)main:goroutine実行時に&wg(WaitGroup)を引き渡す
(5)main:wg.Wait()・・・関数からDone()の通知が来るのを待つ処理をmainの最後に追加する。
*/

func f1(wg *sync.WaitGroup) { // WaitGroupをポインタ引き渡しする
	fmt.Println("f1 (goroutine) execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("f1, i=", i)
		time.Sleep(time.Second)
	}
	fmt.Println("f1 execution finished")
	wg.Done() // goroutineが終了したことをWaitGroupに通知する
	// (*wg).Done()  // この記載でもok (ポインタの説明を参照)
}

func f2() {
	fmt.Println("f1 (goroutine) execution started")
	for i := 5; i < 8; i++ {
		fmt.Println("f2, i=", i)
	}
	fmt.Println("f2 execution finished")
}

func main() {

	fmt.Println("main execution started")

	// 1.WaitGroupの追加
	var wg sync.WaitGroup // 全てのgoroutineが終了するまで待つ
	wg.Add(1)             // 待つgoroutine数 (f1の1つのみ)

	// 2.WaitGroupをgoroutineに引き渡してgoroutine関数実行
	go f1(&wg)
	fmt.Println("No. of Goroutines after go f1():", runtime.NumGoroutine()) //

	f2()

	// 3.Goroutineが全て終わるまで待つ(wg.Wait())
	// time.Sleep(time.Second * 2) // WaitGroupを使うのでこれで待つ必要はないのではずす
	wg.Wait() // goroutineから終了通知(wg.Done())が来るのを待つ

	fmt.Println("main execution stopped")

	/*
		main execution started
		No. of Goroutines after go f1(): 2
		f1 (goroutine) execution started
		f2, i= 5
		f2, i= 6
		f2, i= 7
		f2 execution finished
		f1 (goroutine) execution started
		f1, i= 0
		f1, i= 1
		f1, i= 2
		f1 execution finished
		main execution stopped

		※main開始 => f2 => f1 終了 => main終了
	*/
}
