package main

// naming conventions

var taskDone bool // 4.長いスコープの長い変数名

func main() {

	// 1.予約語 25
	/*
		go specification -> keywords

	*/

	// 2.頭文字を変数名にする
	var mv int // max value

	// 3.短いスコープの変数：長すぎないこと
	var max_value // Python的。GOでは×

	var packetsReceived int // 長すぎる
	var b int // ok

	// 4.長いスコープの変数：長くてok

	// 5.小文字＋大文字 camelCase
	const MAX_VALUE = 100 // not ok

	CONST N = 100 // better

	maxValue := 1000 // recommended
	max_values := 1000 // not ok

	writeToDB := true // ok, idiomatic
	writeToDb := true // not ok
}
