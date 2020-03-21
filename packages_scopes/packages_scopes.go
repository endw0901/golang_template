package main

import (
	"fmt"
)

/*
▼3つのGO scope
(1)File Scope // import
(2)Package Scope
(3)Local Scope(block scope)
*/
const boilingPoint = 100 // 同じパッケージの全てのファイルで使用可能

func sayHello(name string) {
	fmt.Printf("Hello %s!\n", name)
}

func toFahrenheit(t float64) float64 {
	return t*1.8 + 32
}

/*
packages_scopesフォルダのターミナルで、go run *.goと実行すると、
Hello.go、packages_scopes.goの2ファイルが実行される

=> Hello.goファイル内で、packages_scopes.goに定義したsayHello関数をcallできる
*/
