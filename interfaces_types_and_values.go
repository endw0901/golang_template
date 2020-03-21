package main

import (
	"fmt"
	"math"
)

/*
・インターフェイスの初期値はnil
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

	// 0-2.インターフェイス
	print(c1) // shapeインターフェイスを実装するcircle structタイプを受け取るprintメソッド
	print(r1) // shapeインターフェイスを実装するrectangle structタイプを受け取るprintメソッド

	// 1.インターフェイスの初期値:nil
	var s shape
	fmt.Printf("%T\n", s) //<nil>

	// 2.ポリモーフィズム(polymorphism)・・・インターフェイスタイプ(shapeタイプ)の変数(s)は、インターフェイスを実装するnamed typeを保持できる
	ball := circle{radius: 2.5}
	s = ball // shapeインターフェイスを実装するcircle structのballをshapeインターフェイスタイプのsが保持できる

	print(s)
	fmt.Printf("Type of s: %T\n", s) // Type of s: main.circle　※polymorphism

	// 3. runtimeで別の実装タイプにポリモーフィズム(circle struct => rectangle strucタイプ)
	room := rectangle{width: 2, height: 3}
	s = room
	fmt.Printf("Type of s: %T\n", s) // Type of s: main.rectangle
}
