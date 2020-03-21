package main

import "fmt"

func main() {

	// 1.iota
	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)

	fmt.Println(c1, c2, c3)

	// 2.iota省略しても同じ
	const (
		b1 = iota
		b2
		b3
	)

	fmt.Println(b1, b2, b3)

	// 3.便利な使い方
	const (
		North = iota // default 0
		East
		West
		South
	)

	fmt.Println(North, East, West, South)

	// 4.増加2単位
	const (
		e1 = iota * 2
		e2
		e3
	)

	fmt.Println(e1, e2, e3)

	// 5. 開始1から増加2単位
	const (
		g1 = (iota * 2) + 1
		g2
		g3
	)

	fmt.Println(g1, g2, g3)

	// 6.スキップ(_)を使う blank identifier

	// x = -2, y = -4, z = -5

	const (
		x = -(iota + 2) // -(0+2) = -2
		_               // -(1+2) = -3 ※スキップ
		y               // -(2+2) = -4
		z               // -(3+2) = -5
	)

	fmt.Println(x, y, z)

}
