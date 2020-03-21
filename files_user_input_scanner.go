package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*


 */

func main() {

	// 1.実行中二ユーザーの入力を受け取る（ターミナルのコマンドラインで）
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%T\n", scanner) // *bufio.Scanner : ポインター to bufio.Scanner

	scanner.Scan() // 入力待ちとなる

	text := scanner.Text() // string type
	bytes := scanner.Bytes()

	fmt.Println("Input text:", text)   // golang
	fmt.Println("Input bytes:", bytes) // [103 111 108 97 110 103] ※golangのASCIIコード

	// 2.コマンドラインの入力を受け取ってから実行（最初の行のみ）
	/*
		コマンドラインで実行：
		go run files_user_input_scanner.go < files_user_input_scanner.go

		結果：
		*bufio.Scanner
		Input text: package main
		Input bytes: [112 97 99 107 97 103 101 32 109 97 105 110]

		※ファイルの最初の行を返す
	*/

	// 3.全ての入力を受け取る for loop
	for scanner.Scan() {
		text = scanner.Text()
		fmt.Println("You entered:", text)
		if text == "exit" { // exit入力で抜ける。それまでは繰り返す
			fmt.Println("Exiting the scanning...")
			break // for loopを抜けてcontinue
		}
	}

	fmt.Println("Just the message after the for loop")
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
