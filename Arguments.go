package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
コマンドライン引数
command line arguments
*/
func main() {

	// 1.os.Args
	/*
		実行時に引数をつける
		go run Arguments.go python java php
	*/

	fmt.Println("os.Args", os.Args)
	fmt.Println("Path:", os.Args[0])
	fmt.Println("1st argument:", os.Args[1])
	fmt.Println("2nd argument:", os.Args[2])
	fmt.Println("No if items inside os.Args:", len(os.Args))

	// 2.引数に50を指定してもstring
	/*
		実行時に引数をつける
		go run Arguments.go 50 python
	*/
	var result, err = strconv.ParseFloat(os.Args[1], 64) // string->Float
	fmt.Println(result)
	_ = err

	fmt.Printf("%T\n", os.Args[1]) // string
	fmt.Printf("%T\n", result)     // float64
}
