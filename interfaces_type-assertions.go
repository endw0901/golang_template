package main

import (
	"fmt"
	"math"
)

/*
type assertionsはインターフェイスの実装値にアクセスできる
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

// メソッド：
func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
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
	// 0-1.shapeインターフェイスを実装するcircle structタイプをshape変数sで保持する
	var s shape = circle{radius: 2.5}
	fmt.Printf("%T\n", s) // main.circle

	// 1.インターフェイスは実装値にアクセスできない（切り離されている)
	// s.volume() // エラー .\interfaces_type-assertions.go:65:3: s.volume undefined (type shape has no field or method volume)

	// 2.type assertionでインターフェイスの実装値にアクセスする
	ball, ok := s.(circle) // assertion成功：ballがdynamic value保持、失敗：ballは0を保持

	if ok == true {
		fmt.Printf("Ball Volume:%v\n", ball.volume()) // Ball Volume:49.087385212340514
	}

	// 3.switch インターフェイスのtypeでcaseスイッチする

	switch value := s.(type) {
	case circle:
		fmt.Printf("%#v has circle type\n", value)
	case rectangle:
		fmt.Printf("%#v has rectangle type\n", value)
	} // main.circle{radius:2.5} has circle type

	// 4.インターフェイスなので実装タイプを保持できる -> swich case
	s = rectangle{width: 3.4, height: 2.2}

	switch value := s.(type) {
	case circle:
		fmt.Printf("%#v has circle type\n", value)
	case rectangle:
		fmt.Printf("%#v has rectangle type\n", value)
	} // main.rectangle{width:3.4, height:2.2} has rectangle type

}
