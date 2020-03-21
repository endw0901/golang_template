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

	// 0.新規ファイル作成(b.txt) -> 文字を書き込む file.Write
	file, err := os.Open("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// 1.1行ずつファイルスキャン：
	/*
		(1)スキャナー作成：bufio.NewScanner(file)
		(2)スキャナーでスキャン : xxx.Scan()
	*/

	scanner := bufio.NewScanner(file)

	/*
		1) Split指定なし : ファイルそのまま
		2) Split(bufio.ScanWords) : Wordのまとまりごとに表示
		3) Split(bufio.ScanRunes) : Runeごと(1文字ごと)に表示
	*/

	// scanner.Split(bufio.ScanWords) // ファイルをwordごとにスキャン
	// scanner.Split(bufio.ScanRunes) // ファイルをruneごとにスキャン

	success := scanner.Scan() // ポインターを返す
	if success == false {
		err = scanner.Err()
		if err == nil {
			log.Println("Scan was completed and it reached EOF")
		} else {
			log.Fatal(err)
		}
	}

	// 2.スキャンした最初の1行を表示
	// fmt.Println("First line found :", scanner.Text())
	fmt.Println("First line found :", scanner.Bytes())

	// 3.スキャンした全行表示 for loop
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//4. Split 区切り

	/*
		1) Split指定なし : ファイルそのまま
		2) Split(bufio.ScanWords) : Wordのまとまりごとに表示
		3) Split(bufio.ScanRunes) : Runeごと(1文字ごと)に表示
	*/

}
