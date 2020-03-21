package main

import "fmt"

/*
memo
*/

func main() {

	// 1.appendは効率的ではない？capacityが飛び飛び。要素とイコールではない 2->4->8->16
	var nums []int
	fmt.Printf("%#v\n", nums)
	fmt.Printf("Length: %d,Capacity %d \n", len(nums), cap(nums))

	nums = append(nums, 1, 2)
	fmt.Printf("Length: %d,Capacity %d \n", len(nums), cap(nums))

	nums = append(nums, 3)
	fmt.Printf("Length: %d,Capacity %d \n", len(nums), cap(nums)) // Length: 3,Capacity 4 ※3でなく4?

	nums = append(nums, 4)
	fmt.Printf("Length: %d,Capacity %d \n", len(nums), cap(nums)) // cap 4

	nums = append(nums, 100)
	fmt.Printf("Length: %d,Capacity %d \n", len(nums), cap(nums)) // cap 8

	nums = append(nums[0:4], 200, 300, 400, 500, 600) // [1 2 3 4 200 300 400 500 600] 1個上書きして4つ追加(5+4=9でcap8越えるので16へ)
	fmt.Println(nums)
	fmt.Printf("Length: %d,Capacity %d \n", len(nums), cap(nums)) // len 9 , cap 16

	// 2.out of range要素にはアクセスできないが、slice expressionで全てのbacking arrayにアクセスが可能
	letters := []string{"A", "B", "C", "D", "E", "F"}
	letters = append(letters[0:1], "X", "Y")
	fmt.Println(letters)                                                // [A X Y]
	fmt.Printf("Length: %d,Capacity %d \n", len(letters), cap(letters)) // Length: 3,Capacity 6 もともとのA～Fと同じbacking array

	// fmt.Println(letters[4]) // out of range エラー
	fmt.Println(letters[3:6]) // [D E F]
}
