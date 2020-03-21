package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

/*


 */

func main() {

	// 1.新規ファイル作成(b.txt) -> 文字を書き込む file.Write
	file, err := os.Open("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	byteSlice := make([]byte, 2) //2byteを読み込みたいので、2byteのsliceを作成

	// 2. ファイルを指定バイト数読み込む : io.ReadFull(file, slice of byte)
	numberBytesRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Number of bytes read: %d\n", numberBytesRead) //  Number of bytes read: 2
	log.Printf("Data read: %d\n", byteSlice)                  // Data read: [72 101]
	log.Printf("Data read: %s\n", byteSlice)                  // Data read: He

	// 3.全バイト読み込み：ioutil.ReadAll(file)
	// ioutil.ReadAll() reads every byte from the file and return a slice of unknown size

	file, err = os.Open("files_read_bytes.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)

	fmt.Printf("Data as string: %s\n", data)
	fmt.Printf("Numbers of bytes read: %d\n", len(data)) // Numbers of bytes read: 1162

	// 4.全てのファイル内容を素早くbyte sliceに書き出したいときに便利 : ioutil.ReadFile("ファイルパス")
	//** READING WHOLE FILE INTO MEMORY USING ioutil.ReadFile() **//

	// ioutil.ReadFile() reads a file into byte slice
	// this function handles opening and closing the file.
	data, err = ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data read: %s\n", data)

}
