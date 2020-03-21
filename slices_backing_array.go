package main

import (
	"fmt"
	"unsafe"
)

/*
・sliceを作ると、裏でbacking arrayが作られる
・Backing Array(store elements , not slice)
・data structure = slice header

slice header
・address of backing array
len()
cap() capacity

【重要】slice expressionにリターンされたsliceは、オリジナルsliceのbacking arrayを参照する
nil sliceはbacking arrayを持たない
*/

func main() {
	// 1.slice expressionでコピーしたsliceの修正がオリジナルに影響する例
	s1 := []int{10, 20, 30, 40, 50}
	s3, s4 := s1[0:2], s1[1:3] // index 0,1 of s1

	s3[1] = 600 // 10,20 -> 10,600 ※同じbacking arrayを修正することになる

	fmt.Println(s1) // [10 600 30 40 50] ※上書きされた
	fmt.Println(s3) // [10 600]
	fmt.Println(s4) // [600 30] 20,30 -> 600,30 ※上書きされた

	// 2.slice expressionでコピーしたあと、オリジナルを修正した場合 → コピー先も修正される
	arr1 := [4]int{10, 20, 30, 40}
	slice1, slice2 := arr1[0:2], arr1[1:3]

	arr1[1] = 2

	fmt.Println(arr1, slice1, slice2) // [10 2 30 40] [10 2] [2 30]

	// 3.完全に新規にコピーする方法:append
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...) // backing arrayをシェアしない

	cars[0] = "Nissan"         // newCarsは修正されない。
	fmt.Println(cars, newCars) // [Nissan Honda Audi Range Rover] [Ford Honda]

	// 4.cap(), len()
	s5 := []int{10, 20, 30, 40, 50}
	newSlice := s5[0:3]
	fmt.Println(len(newSlice), cap(newSlice)) // 3,5  len=3, capacity=5
	/**
	capacityはorijinalのbacking arrayの要素数
	**/

	newSlice = s5[2:5]                        // 30, 40, 50
	fmt.Println(len(newSlice), cap(newSlice)) // 3,3
	fmt.Println(len(s5), cap(s5))             // 5,5

	// 5. backing arrayが同じでも、slice headerのポインター、アドレスは異なる
	fmt.Printf("%p\n", &s5)                // アドレス (&が必要)
	fmt.Printf("%p,%p \n", &s5, &newSlice) // slice headerは異なるアドレスだが、backing arrayは同じ

	newSlice[0] = 1000
	fmt.Println("s5: ", s5) // s5:  [10 20 1000 40 50]  ※newSlice[0](=s5[2])が変わった ※同じbacking arrayのため

	// 6.arrayとsliceのメモリサイズの比較 ※unsafeパッケージ
	a := [5]int{1, 2, 3, 4, 5} //array 40
	s := []int{1, 2, 3, 4, 5}  //slice 24
	fmt.Printf("array's size in bytes: %d \n", unsafe.Sizeof(a))
	fmt.Printf("slices's size in bytes: %d \n", unsafe.Sizeof(s))
}
