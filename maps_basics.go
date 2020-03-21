package main

import "fmt"

/*
Maps in Go
・array , slice同様に複数要素をstore
・key-valueペアをstore ※dictionary in Python , object in JavaScript
・高速検索 (hash tableを提供するため）
・key同士は同じtype、value同士は同じtypeである必要あり
・ユニークキーを持つ必要あり

・keyはfloatなど、int以外のtypeでもok ※floatは比較に問題があるため推奨されない

・map同士の比較は不可。nilと比較は可能
・mapは順不同
*/

func main() {
	// 1.mapの宣言(初期化無し：nil)
	var employees map[string]string                // key:string value:string
	fmt.Printf("%#v\n", employees)                 // map[string]string(nil)
	fmt.Printf("No of Pairs %d\n", len(employees)) // No of Pairs 0

	// 2.存在しないkeyでmapを検索 -> リターン"" (stringタイプのvalueの初期値は空白)
	fmt.Printf("The value for key %q is %q\n", "Dan", employees["Dan"]) // The value for key "Dan" is ""

	// 3.存在しないkeyでmapを検索 -> リターン0 (float64タイプのvalueの初期値は0)
	var accounts map[string]float64        // key:string, value:float64
	fmt.Printf("%#v\n", accounts["0x323"]) // 0

	// 4.比較できないタイプ(slices)をkeyにもつmap
	/**
	  mapのkeyは比較可能なタイプとする
	  〇int, string, array
	  ×float64
	  ×slices ※複数あるため
	  **/

	// var m1 map[[]int]string // slicesはkeyに使えない
	var m1 map[[5]int]string // arrayはmapのkeyに使える
	_ = m1

	// 5.mapに要素を設定
	// employees["Dan"] = "Programmer" // 初期化せずに設定するとエラーとなる

	people := map[string]float64{}
	people["John"] = 21.4
	people["Mary"] = 10
	fmt.Println(people) // map[John:21.4 Mary:10]

	// 6.初期化済みの空map
	map1 := make(map[int]int)
	map1[4] = 8

	balances := map[string]float64{
		"USD": 323.11,
		"EUR": 432.11,
		// 50:321.11,  // エラー。タイプ違い
		"CHF": 3243.1, // 最後のカンマは必須
	}

	fmt.Println(balances) // map[CHF:3243.1 EUR:432.11 USD:323.11]

	m := map[int]int{1: 10, 2: 20, 3: 30} // 最後のカンマは必須ではない（改行で宣言していないので
	_ = m

	// 7.存在しないのmap keyに設定
	balances["USD"] = 500.2 // -> 更新される
	balances["GBP"] = 300.0 //存在しないkey -> 追加される
	fmt.Println(balances)   // map[CHF:3243.1 EUR:432.11 GBP:300 USD:500.2]

	// 8.存在しないkeyで検索 -> 0
	fmt.Println(balances["RON"]) // 0

	// 9.変数2つにMAP -> 値と、存在するかのboolean
	balances["RON"] = 0
	v, ok := balances["RON"] //

	if ok {
		fmt.Println("The RON balance is: ", v)
	} else {
		fmt.Println("The RON key doesn't exist in map")
	}

	// 10.iterationでmapを出力
	for k, v := range balances {
		fmt.Printf("key: %#v, Value: %#v\n", k, v)
	}

	/*
		key: "USD", Value: 500.2
		key: "EUR", Value: 432.11
		key: "CHF", Value: 3243.1
		key: "GBP", Value: 300
		key: "RON", Value: 0
	*/

	// 11.map要素の削除delete
	delete(balances, "USD")
	fmt.Println(balances) // map[CHF:3243.1 EUR:432.11 GBP:300 RON:0] ※USD削除済み

}
