package main

import "fmt"

/*
map headerへmapの要素へのアドレス・ポインターをメモリに作成する
mapをコピーしても要素はコピーされず、ポインタのコピーとなる
*/

func main() {
	// 1.コピー元を修正すると、コピー先も修正される(ポインタのコピーなので)
	friends := map[string]int{"Dan": 40, "Maria": 25}
	neighbors := friends // 同じ map headerを共有している
	friends["Dan"] = 50
	fmt.Println(neighbors) // map[Dan:50 Maria:25]

	// 2.別々のmapとしてコピーする方法：初期化して、for loopで要素を設定する
	people := make(map[string]int)

	for k, v := range friends {
		people[k] = v // 異なるmap headerを持つ
	}

	people["Anne"] = 18                                       // friendsには影響しない
	fmt.Printf("%#v ---------------- %#v\n", people, friends) // map[string]int{"Anne":18, "Dan":50, "Maria":25} ---------------- map[string]int{"Dan":50, "Maria":25}

	delete(friends, "Dan")
	fmt.Println(friends) // Dan,Maria -> map[Maria:25]

	// 3.keyが存在しなくてもdeleteでエラーは起きない（ヒットしないだけ)
	delete(people, "Lexa")
}
