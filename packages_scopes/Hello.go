package main

import (
	"fmt"
	f "fmt"　// alias"f"つき　=> エイリアスを使えば同じScopeでも重複で使える
)



// ▼同名エラー
// func sayHello(name string) {
// 	fmt.Printf("Hello %s!\n", name)
// }
/*
「packages_scopes.go」で定義済みの関数と同名の関数を定義 => エラー

error:
sayHello redeclared in this block
previous declaration at .\Hello.go:5:20
*/

func main() {
	// 1.別のファイル「packages_scopes.go」で定義した関数をcall
	fmt.Println("Scope means name visibility")
	sayHello("Tanaka")

	// 2.別のファイルで定義したconstantを使って別のファイルで定義した関数をcall
	tf := toFahrenheit(boilingPoint)
	fmt.Println("Water Boiling Point in Degrees Fahrenheit:", tf) // Water Boiling Point in Degrees Fahrenheit: 212

	// 3.エイリアスを使用して同じScopeで重複名のpackageを使う
	f.Println("Something")
}
