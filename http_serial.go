package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

func checkAndSaveBody(url string) {
	resp, err := http.Get(url) // "net/http"パッケージ

	if err != nil {
		fmt.Println(err)
		fmt.Println("%s is DOWN!\n", url)
	} else {
		defer resp.Body.Close()                                      // mainの終了直前までCloseするのを待つ(=deferする)
		fmt.Printf("%s -> Status Code: %d \n", url, resp.StatusCode) // 200：ok、403 Forbidden、404 Not Found
		if resp.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(resp.Body) // byte sliceでbodyを読む
			file := strings.Split(url, "//")[1]         // http://www.google.com  セパレーター //で区切り、0,1の1を採るので「www.google.com」を抽出
			file += ".txt"                              // 「www.google.com.txt」

			fmt.Printf("Writing response body to %s\n", file)

			err = ioutil.WriteFile(file, bodyBytes, 0664) // permission 0664
			if err != nil {
				log.Fatal(err)
			}

		}

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
	urls := []string{"https://golang1.org", "https://www.google.com/a.html", "https://www.medium.com"}

	for _, url := range urls {
		checkAndSaveBody(url)
		fmt.Println(strings.Repeat("#", 20))
	}
}
