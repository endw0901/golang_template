package main

import (
	"fmt"
	"math"
)

/*
・GOには継承（inheritance)の機能がない
 => 複数のinterfaceをマージすることが可能 = interface embedding

*/

// インターフェイス：幾何学形(円・四角)の演算を行う（面積）
type shape interface {
	// circleタイプとrectangleタイプがshapeインターフェイスを実装する
	area() float64
}

type object interface {
	volume() float64
}

// shapeとobjectインターフェイスをembedする(インターフェイスは他のインターフェイスを保持できる)
type geometry interface {
	shape
	object
	getColor() string
}

// geometryにembedされたshapeとobjectのメソッド(area,volume)を全て実装するcube type
type cube struct {
	edge  float64
	color string
}

func (c cube) area() float64 {
	return 6 * (c.edge * c.edge)
}

func (c cube) volume() float64 {
	return math.Pow(c.edge, 3)
}

func (c cube) getColor() string {
	return c.color
}

//
func measure(g geometry) (float64, float64) {
	a := g.area()
	v := g.volume()
	return a, v
}

func main() {

	c := cube{edge: 2}
	a, v := measure(c)
	fmt.Printf("Area:%v, Volume: %v\n", a, v) // Area:24, Volume: 8
}
