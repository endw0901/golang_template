package main

import "fmt"

type km float64
type mile float64

func main() {

	// 1. 基底タイプが同じ underlying type
	type age int
	type oldAge age // intが基底タイプ underlying type
	type veryOldAge age

	// 2.新しいtypeを作る
	type speed uint

	var s1 speed = 10
	var s2 speed = 20

	fmt.Println(s2 - s1) // 同じtype同士の演算は可能

	// 3.タイプ違いエラー（基底タイプに戻して使う
	var x uint
	// x = s1  // s1はtype:speed, xはtype:uint タイプ違いエラー
	x = uint(s1)
	fmt.Println(x)

	// 4.新しいタイプに変換
	var s3 speed = speed(x)
	_ = s3

	// 5.基底タイプの同じ異なる新しい二つのタイプはタイプ違いエラー
	var parisToLondon km = 465
	var distanceInMiles mile

	// distanceInMiles = parisToLondon / 0.621 //タイプ違いエラー(km ≠ mile)
	distanceInMiles = mile(parisToLondon) / 0.621
	fmt.Println(distanceInMiles)
}
