package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// 1.エラーリターン設定→エラー結果判定
	// i, err := strconv.Atoi("45") // ASCII to int , errorをreturn
	i, err := strconv.Atoi("45a") // ASCII to int , errorをreturn

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	// 2.省略形(エラーリターンとエラー結果判定をまとめる)
	if i, err := strconv.Atoi("20a"); err == nil {
		fmt.Println("No error , i is:", i)
	} else {
		fmt.Println(err)
	}

	// 3.引数の数チェック、引数のタイプチェック(数値以外はエラー)。数値以外はエラーだが、引数は文字列string扱い
	if args := os.Args; len(args) != 2 {
		// Pathの他にもう1つ引数が必要(Pathが1個、もう一個)
		// go run simple_if_statement.go xxxx(引数もう一個)
		fmt.Println("One argument is required!")
	} else if km, err := strconv.Atoi(args[1]); err != nil {
		// error handle
		// go run simple_if_statement.go 4a
		fmt.Println("The argument must be an integer! Error:", err)
	} else {
		// go run simple_if_statement.go 4
		fmt.Printf("%d km in miles is %v\n", km, float64(km)*0.621)
		fmt.Printf("%T\n", args[1]) //0:Path , 1:4(4km) -> string
	}

}
