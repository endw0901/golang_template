package main

import (
	"fmt"
	"sync"
	"time"
)

/*
▼data race (concurrency_data_race.go)
data race(race condition) ・・・concurrent systemで最もdebugが困難

例)共有する1つの値について、数百のgoroutineが起動する場合
		→複数のgoroutineでシェアされる変数nの最終値は予測できない(unpredictable)
		=> data raceと呼ばれる
		=> data race detector (コマンドラインで-race)によりチェックが可能。
		=> スレッドセーフ(thread safe)コードにするには？
		=> (1)Mutexes, (2)channels

▼Mutexとは・・・mutual exclusion(相互に排他)
明示的なシンクロ(explicit synchronization)
  => 変数へのアクセスが保護される
　=> 複数のgoroutineでシェアされる変数は、1つの処理が完了するまで他のgoroutineでは読み込まれないように排他的に制御する
※順番を制御するのではなく、アクセスを相互排他にするのみ

*/

func main() {

	// 1.WaitGroup 200作成
	const gr = 100 // goroutineをfor loopで実行する数の分だけwaitgroupを作るのでコンスタントを使う

	var wg sync.WaitGroup
	wg.Add(gr * 2)

	// 2.テスト用の変数nの初期化
	var n int = 0

	// M(1).Mutex text valueを宣言(syncパッケージ)
	var m sync.Mutex

	// 3.200(100ループ中に2つずつ)のgoroutine実行し、n++ n--をそれぞれ行っていく
	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10) // 100ミリ秒(0.1秒)ウェイト

			// M(2) 変数ロック => アンロック ※ロックアンロックは1goroutineにつき1つのみ
			m.Lock() // ロック
			n++
			m.Unlock() // アンロック

			wg.Done()
		}() // anonymous function

		go func() {
			time.Sleep(time.Second / 10) // 100ミリ秒(0.1秒)ウェイト
			// M(2) 変数ロック => アンロック
			m.Lock()
			defer m.Unlock() // 関数を抜ける直前にcallされる。この方法でもok
			n--
			// m.Unlock()

			wg.Done()
		}()

	}

	// 4.全てのgoroutineが終わるまで待つ
	wg.Wait()

	// 5.テスト用変数nの値を出力・・・100回プラス、100回マイナスするので初期値の0になるはず？？
	fmt.Println("Final value of n:", n)

	/*
		▼Mutex付きでの実行結果 => race detectorはdata raceを検知せず、結果も+100-100=0となった
			$ go run -race concurrency_Mutexes.go
			Final value of n: 0
	*/

}
