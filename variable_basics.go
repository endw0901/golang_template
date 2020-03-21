package main

import "fmt"

// 実行・・・ctrl + F5
// 変数

// ERROR -> := is not allowed in package scope (only local scope)
// version := "3.1"

func main() {

	// 1.varを使う / タイプを推測(int等)
	// int age = 10; // c++ way
	var age = 30
	fmt.Println("Age:", age)

	var name = "Dan"
	fmt.Println("your name is:", name)

	var i = 5   // type int
	var j = 5.6 // type float64

	// 2. :=
	s := "Learning golang!"
	fmt.Println(s)

	// 3. := は新規変数のみ
	// name := "Andrei" はエラー
	name = "Andrei"
	name1 := "John"

	// 4.使わない変数はエラーとなるが、_ を使うとエラー回避可能(エラーをミュートする)
	_ = name1

	// 5.複数の変数定義（１）
	car, cost := "Audi", 50000
	fmt.Println(car, cost)

	// 6.1つでも新規変数があればok
	car, year := "Audi", 6000
	fmt.Println(car, year)

	var opened = false
	opened, file := true, "a.txt"

	_, _ = opened, file

	// 7. 複数の変数定義（２）
	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender)

	var a, b, c int
	fmt.Println(a, b, c)

	// 7. 変数の入れ替え
	var i, j int
	i, j = 5, 8
	j, i = i, j
	fmt.Println(i, j)

	// 8.算出
	sum := 5 + 2.3
	fmt.Println(sum)

	// 9.ダブルクオーテーションのみ
	// ERROR -> A string is initialized only using double quotes ("")
	// name = 'Golang'
	name := "Golang" // :=はlocalのみ。packageスコープではエラー。varを使うこと
	fmt.Println(name)

}
