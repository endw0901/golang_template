package main

import (
	"fmt"

	// 1.基本
	// import "fmt" // error

	// 2.エイリアスはok
	f "fmt"
)

// 3.constant(package scope)
const done = false // package scoped
var b int = 10     // package scopeの変数は、未使用でもエラーにならない(localで未使用はエラーとなる)

func main() {
	var task = "Running" // block scope (local)
	fmt.Println(task, done)

	const done = true // local scoped constant -> package scopeのconstantをこのmain内では上書き
	f.Printf("done in main() is %v\n", done)
	f1() //f1内ではpackage scopeのconstant

	f.Println("Bye Bye!")
}

func f1() {
	fmt.Printf("done in f1() is %v\n", done)

	const done = true
	fmt.Printf("done in f1() is %v\n", done)

	a := 10
	_ = a
}
