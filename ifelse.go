package main

import "fmt"

func main() {
	price, inStock := 100, true

	// 1.基本 ※条件文に括弧不要
	if price > 80 {
		fmt.Println("Too Expensive!")
	}

	// _ = inStock

	// 2.boolean省略可能(inStock == trueを省略)
	if price <= 100 && inStock {
		fmt.Println("Buy it!")
	}

	// 3.priceはbooleanではないのでエラー
	//if price {
	//	fmt.Println("something")
	//}

	// 4.ifelse
	if price < 100 {
		fmt.Println("It's Cheap")
	} else if price == 100 {
		fmt.Println("On the edge")
	} else {
		fmt.Println("It's Expensive!")
	}

	// 5.

	age := 7

	if age >= 0 && age < 18 {
		fmt.Printf("You cannot vote! Please return in %d years!\n", 18-age)
	} else if age == 18 {
		fmt.Printf("This is your first vote!")
	} else if age > 18 && age <= 100 {
		fmt.Printf("Please vote, it is important!!")
	} else {
		fmt.Printf("Invalid age")
	}

}
