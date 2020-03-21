package main

import (
	"fmt"
	"log"
	"os"
)

/*
io
ioutil (io - utility)
bufio (buf - io)
*/

func main() {

	// 1.新規ファイル作成
	var newFile *os.File
	fmt.Printf("%T\n", newFile) // *os.File : newFileのタイプは os.Fileへのポインター（メモリアドレス)

	// 2.ファイル作成エラーハンドリング
	var err error
	newFile, err = os.Create("a.txt") //os.Create : 新規ファイル作成。ファイルが存在しているときはtruncate

	/*
		os.Create("a.txt")で、MASTER_GO_PROGRAMMINGフォルダ直下にテキストファイルが生成される
	*/

	if err != nil {
		// fmt.Println(err)
		// os.Exit(1)  // エラーでexitする別の方法

		log.Fatal(err)
		/*
			意図的にエラーを出すには、windows環境だが下記のようなLinux Pathで実行するとエラーテストが可能
			newFile, err = os.Create("/etc/a.txt")

			2020/03/14 13:00:21 open /etc/a.txt: The system cannot find the path specified. ※タイムスタンプ
			exit status 1
		*/
	}

	// 3.ファイルtruncate
	// err = os.Truncate("a.txt",50) // 50byte以内にtruncateする。50byte超過分はロストする
	err = os.Truncate("a.txt", 0) // 0:完全にファイルを空にする ※a.txtに文字を書き込んでから実行するとテキストが空になる事が分かる

	if err != nil {
		log.Fatal(err)
	}

	// 4.ファイルを閉じる
	newFile.Close()

	// 5.ファイルが存在していなければCREATE、存在していればAPPENDモードで開く
	file, err := os.OpenFile("a.txt", os.O_CREATE|os.O_APPEND, 0644) // 存在していなければCREATE、存在していればAPPENDモードで開く
	if err != nil {
		log.Fatal(err)
	}

	file.Close()

	// 6.ファイルの情報を取得する
	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("a.txt")
	p := fmt.Println // 関数省略

	p("File Name:", fileInfo.Name())        // ファイル名を取得 File Name: a.txt
	p("Size in bytes:", fileInfo.Size())    // 0
	p("Last modified:", fileInfo.ModTime()) // 2020-03-14 13:29:40.311602 +0900 JST
	p("Is Directory?:", fileInfo.IsDir())   // false
	p("Permissions?", fileInfo.Mode())      //  -rw-rw-rw-

	if err != nil {
		log.Fatal(err)
	}

	// 7.ファイルが存在するかどうかを確認する

	fileInfo, err = os.Stat("b.txt")

	if err != nil {
		if os.IsNotExist(err) {
			// log.Fatal("File does not exist!") // 2020/03/14 13:32:50 File does not exist!  exit status 1
			fmt.Println("File does not exist!") // 後続処理のためFatalでなくPrintlnにしておく
		}
	}

	// 8.リネーム、ファイルの移動
	oldPath := "aaa.txt"
	newPath := "aa.txt"
	err = os.Rename(oldPath, newPath) // a.txt -> aaa.txtにリネームされる

	if err != nil {
		log.Fatal(err)
	}

	// 9.ファイル削除
	err = os.Remove("aa.txt")
	if err != nil {
		log.Fatal(err)
	}

}
