package main

import "fmt"

/*
UTF8

*/

func main() {
	s1 := "Hi there Go!"
	fmt.Println(s1)

	// 1.エスケープ (1)バックスラッシュ、(2)バックティック
	fmt.Println("He says: \"Hello\"")
	fmt.Println(`He says: "Hello!"`)

	// 2.バックティックは表示されない。バックティック内の\nは改行ではなく文字扱い
	s2 := `I like \n Go`
	fmt.Println(s2) // I like \n Go

	// 3.バックティックで改行
	fmt.Println("Price: 100\n Brand: Nike")
	fmt.Println(`Price: 100\n Brand: Nike`)
	fmt.Println(`
	Price: 100
	 Brand: Nike
	 `)

	// 4.アドレスにはバックスラッシュのエスケープが必要だがバックスラッシュなら個別エスケープが不要
	fmt.Println(`C:\Users\Andrei`)
	fmt.Println("C:\\Users\\Andrei")

	// 5.コンカチ
	var s3 = "I love " + "Go " + "Programming"
	fmt.Println(s3 + "!")

	// 6.要素
	fmt.Println("Element at index 0:", s3[0]) // 73:最初の文字「I」のASCIIコード

	// 7.引用
	fmt.Printf("%s\n", s3) // I love Go Programming
	fmt.Printf("%q\n", s3) // "I love Go Programming"

}
