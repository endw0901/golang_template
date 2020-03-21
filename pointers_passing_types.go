package main

import "fmt"

/*
▼基本
   *float64 : タイプ宣言
   *p : ポインター変数

   &value   => pointer
   *pointer => value

   **pointer => value(pointer)のvalue

▼関数内で引き渡したタイプを呼び出し元(main)で変更するのにポインタ引渡しが必要なタイプ
必要：int, float64, string, bool, struct　※変数そのままの引渡しではmain上は変更されない
不要：slice, map ※ポインタを使わなくても、そのまま引き渡せば関数内の変更が呼び出し元(main)上のslice,mapに反映される

※ベストプラクティスでは、arrayではなくsliceを関数に引き渡す

▼注意事項
ポインターの引渡しは、garbage collectionに負荷をかけることになるため、
必要な場合以外は使わないこと

*/

// 関数内の変数の変更のみ
func changeValues(quantity int, price float64, name string, sold bool) {
	quantity = 3
	price = 500.4
	name = "Mobile phone"
	sold = false
}

// 呼び出し元の入力変数の値を変更する
func changeValuesByPointer(quantity *int, price *float64, name *string, sold *bool) {
	/*
	 *int:タイプ宣言
	 *price:ポインター変数に格納されたポインターが指す変数に格納された値
	 */
	*quantity = 3
	*price = 500.4
	*name = "Mobile phone"
	*sold = false
}

// structタイプ宣言
type Product struct {
	price       float64
	productName string
}

// 関数内のstructの項目変更のみ
func changeProduct(p Product) {
	p.price = 300
	p.productName = "Bicycle"
}

// 呼び出し元(main)の入力structの項目変更
func changeProductByPointer(p *Product) {
	(*p).price = 300 // p.price=300としても同じ
	p.productName = "Bicycle"
}

// 入力sliceを関数内で変更
func changeSlice(s []int) {
	for i := range s {
		s[i]++
	}
}

// 入力mapを関数内で変更
func changeMap(m map[string]int) { // keyがstring, valueがintのmap
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30
}

func main() {

	// 1.関数に引き渡した変数の関数内の修正は、呼び出し元(main)の変数の値を変更しない
	quantity, price, name, sold := 5, 300.4, "Laptop", true
	fmt.Println("Before calling changeValues():", quantity, price, name, sold) // 5 300.4 Laptop true
	changeValues(quantity, price, name, sold)
	fmt.Println("After calling changeValues():", quantity, price, name, sold) // 5 300.4 Laptop true

	// 2.関数内で、呼び出し元(main)の変数を変更したい場合　=> pointerを使う
	fmt.Println("Before calling changeValuesByPointer():", quantity, price, name, sold) // 5 300.4 Laptop true
	changeValuesByPointer(&quantity, &price, &name, &sold)
	fmt.Println("After calling changeValuesByPointer():", quantity, price, name, sold) // 3 500.4 Mobile phone false

	// 3.structを関数に引き渡す =>変更されない

	gift := Product{
		price:       100,
		productName: "Watch",
	}

	fmt.Println(gift) // {100 Watch}
	changeProduct(gift)
	fmt.Println(gift) // {100 Watch}

	// 4.structのポインタを関数に引き渡し、関数内でmainのstruct項目を変更する
	fmt.Println(gift) // {100 Watch}
	changeProductByPointer(&gift)
	fmt.Println(gift) // {300 Bicycle}

	// 5.slice,mapを関数に引渡し、関数内で変更 => mainのsliceが変更される (ポインタ不要)
	prices := []int{1, 2, 3}
	fmt.Println("prices slice before calling changeSlice():", prices) //[1 2 3]
	changeSlice(prices)
	fmt.Println("prices slice after calling changeSlice():", prices) // [2 3 4]

	myMap := map[string]int{"a": 1, "b": 2}
	fmt.Println("prices slice before calling changeMap():", myMap) // map[a:1 b:2]
	changeMap(myMap)
	fmt.Println("prices slice after calling changeMap():", myMap) //  map[a:10 b:20 c:30]

}
