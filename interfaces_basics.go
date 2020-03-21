package main

import (
	"fmt"
	"math"
)

/*
▼インターフェイス
・メソッドの集まり
・多くはnamed type
・Javaのようなgeneric typeではない

・インターフェイスはメソッドのsignature(入力変数とリターン値)のみで、implementはしない
・Javaのようにimplementsの明記は不要。named typeが実装していればinterfaceの実装とみなす
  (下記例では、circleタイプとrectangleタイプがshapeインターフェイスを実装する)

・インターフェイスtypeを受け取るメソッドを定義することで、インターフェイスを実装するどのnamed typeも入力として受け取れる
  =>これによって複数のnamed typeのメソッドを共通化できる
*/

// インターフェイス：幾何学形(円・四角)の演算を行う（面積・外周）
type shape interface {
	// circleタイプとrectangleタイプがshapeインターフェイスを実装する
	area() float64
	perimeter() float64
}

// struct:四角形、円
type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

// メソッド：円の面積計算(receiver:circle struct)
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// メソッド：円周計算
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// メソッド：四角形の面積計算(receiver:rectangle struct)
func (r rectangle) area() float64 {
	return r.height * r.width
}

// メソッド：四角形の外周
func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// メソッド：円のメソッドをcallしてprintするメソッド
// func printCircle(c circle) {
// 	fmt.Println("Shape:", c)
// 	fmt.Println("Area:", c.area())
// 	fmt.Println("Perimeter:", c.perimeter())
// }

// メソッド：四角形のメソッドをcallしてprintするメソッド
// func printRectangle(r rectangle) {
// 	fmt.Println("Shape:", r)
// 	fmt.Println("Area:", r.area())
// 	fmt.Println("Perimeter:", r.perimeter())
// }

// メソッド：shapeインターフェイスを受け取る = shapeインターフェイスを実装するどのtypeも入力変数として受け取れる
func print(s shape) {
	fmt.Printf("Shape: %#v\n", s)
	fmt.Printf("Area: %#v\n", s.area())
	fmt.Printf("Perimeter: %#v\n", s.perimeter())
}

func main() {

	// 0-1.structの定義
	c1 := circle{radius: 5.}
	r1 := rectangle{width: 3, height: 2.1}

	// 0-2.メソッドcall
	// printCircle(c1)
	// printRectangle(r1)

	/*
	   同じロジックのメソッドをcircle, rectangleそれぞれ作っていて無駄
	   Shape: {5}
	   Area: 78.53981633974483
	   Perimeter: 31.41592653589793

	   Shape: {3 2.1}
	   Area: 6.300000000000001
	   Perimeter: 10.2
	*/

	// 1.インターフェイスで上記メソッドをリファクタリング
	/* 幾何学形(円・四角)の演算を行う（面積・外周）インターフェイス */

	print(c1) // shapeインターフェイスを実装するcircle structタイプを受け取るprintメソッド
	print(r1) // shapeインターフェイスを実装するrectangle structタイプを受け取るprintメソッド

	/*
		Shape: main.circle{radius:5}
		Area: 78.53981633974483
		Perimeter: 31.41592653589793

		Shape: main.rectangle{width:3, height:2.1}
		Area: 6.300000000000001
		Perimeter: 10.2
	*/
}
