package main

import "fmt"

/*
memo
*/

func main() {

	// 1.x[start:stop]
	a := [5]int{1, 2, 3, 4, 5}
	// a[start:stop] // include start, exclude stop

	b := a[1:3]                  // index 1,2 (exclude 3) →valueは2(index1),3(index2)
	fmt.Printf("%v, %T\n", b, b) // [2 3], []int ※typeはslice

	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[1:3]

	fmt.Println(s2)     // [2 3]
	fmt.Println(s1[2:]) // [3 4 5 6]  same as s1[2:len(s1)]
	fmt.Println(s1[:3]) // [1 2 3]    same as s1[0:3]
	fmt.Println(s1[:])  // [1 2 3 4 5 6] same as s1[0:len(s1)] ※全部
	// fmt.Println(s1[:100]) // out of rangeエラー

	// 2.apend
	s1 = append(s1[:4], 100) //index3の後に100を追加
	fmt.Println(s1)          // [1 2 3 4 100]

	s1 = append(s1[:4], 200) // 100を200に上書き
	fmt.Println(s1)          //[1 2 3 4 200]
}
