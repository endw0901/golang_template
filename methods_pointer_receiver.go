package main

import (
	"fmt"
)

/*
▼メソッド(ポインタ引き渡し）を使う時
(1)関数に渡すと、変数がローカルにコピーされる。そのコピーの分容量を食うのを避けたいとき
(2)引き渡す値を変更したいとき
*/

// method用のtype =>引数となる
type car struct {
	brand string
	price int
}

// 関数
func changeCar(c car, newBrand string, newPrice int) {
	c.price = newPrice
	c.brand = newBrand
}

// メソッド(値を引き渡し)
func (c car) changeCar1(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

// メソッド（ポインタを引き渡し）
func (c *car) changeCar2(newBrand string, newPrice int) {
	(*c).brand = newBrand // *cポインタが指す場所を変える
	(*c).price = newPrice
	// *c.priceではエラーとなる。括弧は必須
}

/*
▼ポインタタイプの引き渡しはエラー　※メソッドはvalue typeでのみ作成可能。ポインタタイプは不可

type distance *int

func (d *distance) m1() {
	fmt.Println("just a message")
}
*/

func main() {

	// 0.関数
	myCar := car{brand: "Audi", price: 40000}
	changeCar(myCar, "Renault", 20000) // 関数はローカルコピーで動くため、mainの変数を変更しない
	fmt.Println(myCar)                 // {Audi 40000}

	// 1.メソッド （値を引き渡し）
	myCar.changeCar1("Opel", 21000)
	fmt.Println(myCar) // {Audi 40000} ※変わらない

	// 2.メソッド（ポインターを引き渡し） => 変更される
	(&myCar).changeCar2("Seat", 25000)
	fmt.Println(myCar) // {Seat 25000}

	// 3.省略形 (&myCar). => myCar.でもok
	myCar.changeCar2("Seat", 25000)
	fmt.Println(myCar) // {Seat 25000}

	// 4. ポインタタイプを使ったメソッドcall
	var yourCar *car
	yourCar = &myCar
	fmt.Println(*yourCar) // {Seat 25000}  *yourCar ※yourCarポインタが示す値を取得するのが「*yourCar」

	yourCar.changeCar2("VW", 30000)
	fmt.Println(*yourCar) // {VW 30000}

	(*yourCar).changeCar2("Another VW", 30000)
	fmt.Println(*yourCar) // {Another VW 30000}

	fmt.Println(myCar) // {Another VW 30000} ※元のmyCarも変更された

}
