package main

import (
	"fmt"
	"sync"
	"time"
)

/*
data race(race condition) ・・・concurrent systemで最もdebugが困難

例)共有する1つの値について、数百のgoroutineが起動する場合
		→複数のgoroutineでシェアされる変数nの最終値は予測できない(unpredictable)
		=> data raceと呼ばれる
		=> スレッドセーフ(thread safe)コードにするには？
		=> concurrency_Mutexes.go参照

▼race detectorについて
Linux, Mac OS, Windows system for 64 bit processorsで提供されている = race detector
c, c++ runtime library

GOにはbuilt in race detectorがある

race detectorは、concurrentプログラムが正しく作られているかの強力なチェックツール
(コマンドラインで実行可能)

 => data raceを避け、スレッドセーフコードを書くには？
 => concurrency_Mutexes.go

▼race detectorの機能
・race flagをコマンドラインツールに立てる
=> コンパイラが、いつどのようにメモリがアクセスされたかのrecordを取るようになる
   かつ、runtimeライブラリがシェアされた変数に対する非同期のアクセスを監視するようになる

▼コマンドラインでrace detectorを実行
go run -race concurrency_data_race.go
*/

func main() {

	// 1.WaitGroup 200作成
	const gr = 100 // goroutineをfor loopで実行する数の分だけwaitgroupを作るのでコンスタントを使う

	var wg sync.WaitGroup
	wg.Add(gr * 2)

	// 2.テスト用の変数nの初期化
	var n int = 0

	// 3.200(100ループ中に2つずつ)のgoroutine実行し、n++ n--をそれぞれ行っていく
	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10) // 100ミリ秒(0.1秒)ウェイト
			n++
			wg.Done()
		}() // anonymous function

		go func() {
			time.Sleep(time.Second / 10) // 100ミリ秒(0.1秒)ウェイト
			n--
			wg.Done()
		}()

	}

	// 4.全てのgoroutineが終わるまで待つ
	wg.Wait()

	// 5.テスト用変数nの値を出力・・・100回プラス、100回マイナスするので初期値の0になるはず？？
	fmt.Println("Final value of n:", n)

	/*
		1回目・・・Final value of n: 3
		2回目・・・Final value of n: -5
		3回目・・・Final value of n: 1

		※処理Aがnを0->1にする間に処理Bが0を-1にするかもしれない。 0->1->-1となり、+1-1=0とはならない。
		→複数のgoroutineでシェアされる変数nの最終値は予測できない(unpredictable)
		=> data raceと呼ばれる
		=> スレッドセーフ(thread safe)コードにするには？
	*/

	/*
		▼ race detector(-raceつき)つきで実行した結果

		==================
		Final value of n: -3
		Found 4 data race(s)
		exit status 66

	*/
}
