package main

import (
	"io/ioutil"
	"log"
	"os"
)

/*
ファイルにバイトを書き込む
・os.Write
・ioutil.WriteFile


io, ioutil, bufioは機能が多いが特に必要ないものも多く、メモリサイズを食うため
osパッケージで書けばよい

*/

func main() {

	// 1.新規ファイル作成(b.txt) -> 文字を書き込む file.Write
	file, err := os.OpenFile(
		"b.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644, // mode
	)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close() // deferは別に詳細を記載する

	byteSlice := []byte("I learn Golang!")     // stringをbyte sliceに変換する -> byte sliceをファイルに変換する /15byte
	bytesWritten, err := file.Write(byteSlice) //書き込んだbyteをリターンする

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Bytes Written: %d\n", bytesWritten) // 2020/03/14 13:51:57 Bytes Written: 15

	// 2.ioutil.WriteFile("新規作成ファイル名",byte-slice, mode) : slice of bytesをすばやくファイルにダンプしたいときに便利

	bs := []byte("Go is good")                // ASCIIコード10byte
	err = ioutil.WriteFile("c.txt", bs, 0644) //ファイルがなければ新規に作成し、既に存在していればtruncateしてから書き込む

	if err != nil {
		log.Fatal(err)
	}

}
