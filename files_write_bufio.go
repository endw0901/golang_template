package main

import (
	"bufio"
	"log"
	"os"
)

/*
ファイルにバイトを書き込む
・os.Write
・ioutil.WriteFile


io, ioutil, bufioは機能が多いが特に必要ないものも多く、メモリサイズを食うため
osパッケージで書けばよい

◆bufio
ファイル書き込みの際、ディスク書き込みの前にメモリにバッファーを持つ。
沢山の操作がファイル書き込み前に必要な時に便利
(ファイル書き込みの速度が速くなる。操作してから書き込むため)

*/

func main() {

	// 1.新規ファイル作成(b.txt) -> 文字を書き込む file.Write
	file, err := os.OpenFile("b.txt", os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// 2. バッファーにslice of byte書き込み : bufferedWriter.Write(slice of byte)
	bufferedWriter := bufio.NewWriter(file)
	bs := []byte{97, 98, 99} // slice of byte
	bytesWritten, err := bufferedWriter.Write(bs)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Bytes written to buffer (not file) %d\n", bytesWritten) // Bytes written to buffer (not file) 3

	// 3.使えるバッファサイズを確認 : 4096 - 3 = 4093残り
	bytesAvailable := bufferedWriter.Available()
	log.Printf("Bytes available in buffer: %d\n", bytesAvailable) // Bytes available in buffer: 4093

	// 4.バッファーにstring書き込み : bufferedWriter.WriteString("\nストリング")
	bytesWritten, err = bufferedWriter.WriteString("\nJust a random string") // 21bytes

	if err != nil {
		log.Fatal(err)
	}

	// 5.バッファーに書き込まれたサイズを確認  : bufferedWriter.Buffered()
	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize) // Bytes buffered: 24 (3byte(slice) + 21bytes(string))

	/*
		ここまではバッファーに書き込まれていて、ファイルには書き込みされていない
		ここからメモリダンプ
	*/

	// 6.バッファーをファイルに書き出し（メモリダンプ）
	bufferedWriter.Flush()

	// 7.メモリに書き込みされていないバッファー変更をリセットしたいとき
	bufferedWriter.Reset(bufferedWriter)

}
