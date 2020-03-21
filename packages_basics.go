package main

import (
	"fmt"
	"mypackages/numbers"
)

/*
▼executable packages
・実行可能ファイルを作成する
・mainと定義される

▼non-executable packages
・どのような名前でも定義できる
・他のパッケージで利用される。
・実行不可で、importされて使用されるのみ

▼環境変数：GOPATH
Windowsでは %USERPROFILE%\go
MINGW64(bash) -> go env

▼
・実行可能なファイルは全てpackage main
・import pathは全て、$GOPATH/src の配下にある
 =>import  mypackages/numbersと指定すると、$GOPATH/src配下を検索する

・packageとdirectoryは同じ名前である必要はない
 => mypackages/numbers directoryのnumbers.goをnumbers1.goにリネームしても、このソースは変える必要がない

 ▼importするpackageの関数名
 ・package内の関数名は大文字で始めまる名前にすること
  => 大文字で始まる関数のみが他のpackageからアクセスができる
*/

func main() {

	var x int = 40

	// 1.mypackages/numbers配下のEven関数を呼び出す(directory配下のソースを読み込む)
	fmt.Printf("%d is even: %t \n", x, numbers.Even(x))

	/*
		numbers.Even()では、mypackages/numbersディレクトリを読み込んでいるのであって、
		mypackages/numbers配下のnumbers1.goファイル名を指定するのではない
	*/

	// 2.mypackages/numbers配下のIsPrime関数を呼び出す
	fmt.Println(numbers.IsPrime(13), numbers.IsPrime(100)) // true false

	// 3.mypackages/numbers配下のpublic関数を呼び出す(関数内でprivate関数を使っている関数)
	fmt.Println(numbers.OddAndPrime(25)) //false
	fmt.Println(numbers.OddAndPrime(13)) //true
	fmt.Println(numbers.OddAndPrime(17)) //true

	// fmt.Println(numbers.isPrime(17)) // private関数のためエラー：cannot refer to unexported name numbers.isPrime

}
