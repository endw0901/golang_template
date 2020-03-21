package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"
)

/*
▼処理内容
URLのリストをチェックする
response bodyをsaveして出力
※web serverとAPIの接続を継続的にチェックする処理に使えるような処理

▼もし1000サイトあるとすると、
1サイト2秒×1000時間がかかってしまう
=>並行でcheckすれば数秒で終わることになる
=> http_serial.go => http_waitgroups.goではconcurrent化
*/

func checkAndSaveBody(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url) // "net/http"パッケージ

	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is DOWN!\n", url)
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

	wg.Done()
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

	// 1.WaitGroupを作成
	var wg sync.WaitGroup

	// 2.WaitGroupに待つgoroutine数を追加
	wg.Add(len(urls)) // checkするurlの数だけgoroutineがあるのでurlのlengthをwaitgroupにaddする

	// 3.check関数にWaitGroupポインタを引き渡しcall => 関数で受け取り、wg.Done()で終了通知
	for _, url := range urls {
		go checkAndSaveBody(url, &wg)
		fmt.Println(strings.Repeat("#", 20))
	}

	fmt.Println("No of Goroutines:", runtime.NumGoroutine())
	// 4.goroutineが終了するまでmainの終了を待つ
	wg.Wait()

	/*
		$ go run http_waitgroups.go
		####################
		####################
		####################
		Get "https://golang1.org": dial tcp: lookup golang1.org: no such host
		%s is DOWN!
		 https://golang1.org
		https://www.google.com/a.html -> Status Code: 404
		https://www.medium.com -> Status Code: 200
		Writing response body to www.medium.com.txt
	*/
}
