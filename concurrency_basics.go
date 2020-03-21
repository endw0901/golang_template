package main

import (
	"fmt"
	"runtime"
)

/*
・Concurrencyをbuilt-inで言語レベルで優先してサポートする初のメジャー言語
・2012年3月にGO ver1.0がリリースされる

(1)ParallelismとConcurrencyの違い
Parallelismは複数CPUが必要。同時に複数のGOルーチンが実行されること
Concurrencyは複数のGoルーチンをloadし、1つが止まったらもう1つが動き始める。1つのCPUならconcurrentアプリのみ稼働する
Concurrencyは複数のGoルーチンをそれぞれ独立したものとして同時に動かせる

(2)Goroutineとは：fixedスレッドとの違い
・数十万スレッドほども稼働できる軽量なスレッド(lightweight thread)
・2キロバイトほどのstack spaceがあればいい(1～2MB必要なfixed threadと比較してもかなり軽量)
2MBは巨大なメモリ占有だが、複雑な再起関数にとっては足りない
・Goroutineは必要に応じて拡張・縮小する(max 1GBほど)

(3)OSスレッドとGoroutine
・OSスレッドはOS kernelによりスケジュールされる
・Goroutineは、独自にm:nスケジューリングを用いる n OSスレッド上にm goroutineをスケジュールする
軽量なスケジューリング
・Goroutineには、プログラマーがアクセスできる識別子identityはない

*/

func f1() {
	fmt.Println("f1 (goroutine) execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("f1, i=", i)
	}
	fmt.Println("f1 execution finished")
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

	// 1.CPU数：NumCPU()・・・現在のプロセスで使用可能なcpu数を表示
	fmt.Println("No. of CPUs :", runtime.NumCPU()) // 4

	// 2.NumGoroutine()・・・現在存在しているGoroutine数
	fmt.Println("No. of Goroutines :", runtime.NumGoroutine()) // 1 (main)

	// 3.GOOS:
	fmt.Println("OS:", runtime.GOOS) // OS: windows

	// 4.アーキテクチャ
	fmt.Println("Arch:", runtime.GOARCH) // Arch: amd64

	// 5.実行中OSスレッドの数　goスケジューラーはGOMAXPROCS パラメーターで 何個のOSスレッドを同時に動かすか決める
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0)) // GOMAXPROCS(0) 0は現在の設定をそのまま返す => GOMAXPROCS: 4

	// 5.複数のgoroutineを実行(main, f1) => f1が終了する前にmainが終了してしまう問題が発生
	go f1()
	fmt.Println("No. of Goroutines after go f1():", runtime.NumGoroutine()) //

	f2()

	// time.Sleep(time.Second * 2) // これを挿入するとf1が終了する時間ができる
	fmt.Println("main execution stopped")

	/*
		No. of Goroutines after go f1(): 2
		f1 (goroutine) execution started
		f2, i= 5
		f2, i= 6
		f2, i= 7
		f2 execution finished
		main execution stopped

		※mainとf1が同時に動く => mainがf1を待たずに終了 => f1のexitメッセージが実行されずに終了
		  => シンクロさせる仕組みが必要 => time.Sleep挿入でf1がexitする時間を作れる
		=> 2秒で終わるとは限らないという問題
		=> 別の方法が必要 => wait group
	*/
}
