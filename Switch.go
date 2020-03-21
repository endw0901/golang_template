package main

import (
	"fmt"
	"time"
)

/*
memo
・読みやすいので使う
*/

func main() {

	// 1.基本:switch 変数
	language := "golang" // caseも文字列。caseで数字を使うとエラーとなる

	switch language {
	case "Python": // 10とかにするとエラー
		fmt.Println("You are learning Python!You don't use curly braces but indentation!!")
		// break //breakは書かなくてok
	case "Go", "golang": // or
		fmt.Println("Good, go fo go! You are using curly braces {}!")

	default: //ifのelseと同等
		fmt.Println("Any other programming language is a good start!")
		/*
			・defaultは省略可能。
			・どのcaseにも合致しない時のみdefaultに入る
		*/
	}

	// 2.switch bool(true/false)
	n := 5

	switch true {
	case n%2 == 0:
		fmt.Println("even number!")
	default:
		fmt.Println("odd number!")
	}

	// 3.最初に条件合致したところで抜ける
	hour := time.Now().Hour() // built-inパッケージtime
	fmt.Println(hour)

	switch { // switch trueは省略可能(switchのみでok)
	case hour < 12:
		fmt.Println("Good Morning!")
	case hour < 17: // 条件は合致するが、上の条件合致でswich文を抜ける
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}
