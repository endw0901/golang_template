package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
)

/*
▼処理内容
URLのリストをチェックする
response bodyをsaveして出力
※web serverとAPIの接続を継続的にチェックする処理に使えるような処理

▼もし1000サイトあるとすると、
1サイト2秒×1000時間がかかってしまう
=>並行でcheckすれば数秒で終わることになる
=> http_waitgroups.goを参照
*/

func checkAndSaveBody(url string, c chan string) {
	/*
		・引数にchannelを追加
		・errリターンの文字列をchannel通信
	*/
	resp, err := http.Get(url) // "net/http"パッケージ

	if err != nil {
		//fmt.Println(err)
		//fmt.Println("%s is DOWN!\n", url)
		s := fmt.Sprintf("%s is DOWN!\n", url)
		s += fmt.Sprintf("Error: %v\n", err) // urlとerrをコンカチ
		c <- s                               // channelに送る
	} else {
		defer resp.Body.Close() // mainの終了直前までCloseするのを待つ(=deferする)
		// fmt.Printf("%s -> Status Code: %d \n", url, resp.StatusCode) // 200：ok、403 Forbidden、404 Not Found
		s := fmt.Sprintf("%s -> Status Code: %d \n", url, resp.StatusCode) // SprintfはPrintfと同じ挙動をするが、Sprinfはstringを返す

		if resp.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(resp.Body) // byte sliceでbodyを読む
			file := strings.Split(url, "//")[1]         // http://www.google.com  セパレーター //で区切り、0,1の1を採るので「www.google.com」を抽出
			file += ".txt"                              // 「www.google.com.txt」

			// fmt.Printf("Writing response body to %s\n", file)
			s := fmt.Sprintf("Writing response body to %s\n", file)

			err = ioutil.WriteFile(file, bodyBytes, 0664) // permission 0664
			if err != nil {
				// log.Fatal(err)
				s += "Error writing file\n"
				c <- s
			}

		}

		s += fmt.Sprintf("%s is UP\n", url) // respがエラーではないときのリターン
		c <- s
	}
}

/*
https://en.wikipedia.org/wiki/List_of_HTTP_status_codes
1xx informational response – the request was received, continuing process
2xx successful – the request was successfully received, understood, and accepted
3xx redirection – further action needs to be taken in order to complete the request
4xx client error – the request contains bad syntax or cannot be fulfilled
5xx server error – the server failed to fulfil an apparently valid request
*/

func main() {

	// 0.checkするURLを指定
	urls := []string{"https://golang1.org", "https://www.google.com/a.html", "https://www.medium.com"}

	// 1.channel宣言
	c := make(chan string)

	// 2.goroutine起動＆channel通信開始
	for _, url := range urls {
		go checkAndSaveBody(url, c) // 関数のリターン文字列をchannelで通信する
		fmt.Println(strings.Repeat("#", 20))
	}

	fmt.Println("Num of goroutines:", runtime.NumGoroutine())

	// 3.channelに蓄積された関数のリターン文字列を、関数の数(=checkするURLの数)だけ出力する
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}

	/*
		####################
		####################
		####################
		Num of goroutines: 6
		https://golang1.org is DOWN!
		Error: Get "https://golang1.org": dial tcp: lookup golang1.org: no such host

		https://www.google.com/a.html -> Status Code: 404
		https://www.google.com/a.html is UP

		https://www.medium.com -> Status Code: 200
		https://www.medium.com is UP
	*/
}
