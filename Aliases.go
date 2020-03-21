package main

import "fmt"

/*
Aliasは普段使わない
・大規模システムの移行などで使う

runeはuint32のalias

alias = 同じタイプの別名
aliasはタイプ違いのエラーは発生しない
*/

func main() {

	// 1.byte, uint8
	var a uint8 = 10
	var b byte //byteはuint8のalias

	b = a
	_ = b

	// 2.aliasの作り方
	type second = uint
	var hour second = 3600
	fmt.Printf("Minutes in an hour: %d \n ", hour/60)

}
