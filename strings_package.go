package main

import (
	"fmt"
	"strings"
)

/*
memo
*/

func main() {
	// 1.よく使うPrintln関数を短縮化
	p := fmt.Println

	// 2.Contains:含むかどうかtrue/false
	result := strings.Contains("I love Go Programming", "love")
	p(result) // true

	result = strings.Contains("I love Go Programming", "lovex")
	p(result) // false

	// 3.ContainsAny:いずれかを含むかどうか xysなら、sがあればok
	result = strings.ContainsAny("success", "xy") // xもyもsuccessにはないのでfalse
	p(result)                                     // false

	result = strings.ContainsAny("success", "xys") // sがsuccessに含まれるのでtrue
	p(result)                                      // true

	// 4.ContainsRune
	result = strings.ContainsRune("golang", 'g') // runeは'' ""ではなく'' シングルクォーテーション
	p(result)

	// 5. Count:文字列に含まれる文字の数
	n := strings.Count("cheese", "e") // eがcheeseに何個含まれるか
	p(n)

	// 6.Countで空検索→rune+1を返す
	n = strings.Count("Five", "") // Fiveの 4 rune + 1 = 5
	p(n)                          // 5

	// 7.ToLower:小文字変換, ToUpper:大文字返還
	p(strings.ToLower("GO PyTHON jaVA")) // go python java
	p(strings.ToUpper("GO PyTHON jaVA"))

	// 8. 効率よく文字列を比較する EqualFold(大文字小文字を無視)
	p("go" == "go") // true

	p("GO" == "go")                                   // false
	p(strings.ToLower("GO") == strings.ToLower("go")) // true works but効率的でない

	p(strings.EqualFold("GO", "go")) // true 効率的(推奨の比較方法)

	// 9.Repeat
	myStr := strings.Repeat("ab", 10)
	p(myStr) // abababababababababab

	// 10.Replace, ReplaceAll
	myStr = strings.Replace("192.168.0.1", ".", ":", 2) // 最初の2つの.を:に置換
	p(myStr)                                            // 192:168:0.1

	myStr = strings.Replace("192.168.0.1", ".", ":", -1) // -1 は全て置換となる
	p(myStr)                                             // 192:168:0:1

	myStr = strings.ReplaceAll("192.168.0.1", ".", ":")
	p(myStr) // 192:168:0:1

	// 11.Split (1)カンマ区切り (2)全文字分割
	s := strings.Split("a,b,c", ",") // 区切り文字 ,
	fmt.Printf("%T\n", s)            // []string  = slice of strings
	fmt.Printf("%#v\n", s)           // []string{"a", "b", "c"}

	s = strings.Split("Go fo Go!", "")
	fmt.Printf("%#v\n", s) // []string{"G", "o", " ", "f", "o", " ", "G", "o", "!"}

	// 12.Join : 結合
	s = []string{"I", "learn", "Golang"}
	myStr = strings.Join(s, "-")
	p(myStr) // I-learn-Golang

	// 13.Field : slice of stringsを返す
	myStr = "Orange Green \n Blue Yello"
	fields := strings.Fields(myStr)
	fmt.Printf("%T\n", fields)  // []string  : slice of strings
	fmt.Printf("%#v\n", fields) // []string{"Orange", "Green", "Blue", "Yello"}

	// 14.TrimSpace : tabや改行を除去(空白は残る)
	s1 := strings.TrimSpace("\t Goodbye Windows, Welcome Linux! \n") //
	fmt.Printf("%q\n", s1)                                           // "Goodbye Windows, Welcome Linux!"

	// 15.Trim : 指定文字を除去
	s2 := strings.Trim("...Hello Gophers!!!!?", ".!?") // .と!と?を除去
	fmt.Printf("%q\n", s2)                             // "Hello Gophers"
}
