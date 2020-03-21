package main

import "fmt"

/*
nested struct (embedded struct)
*/

func main() {
	// 1.structの中にstructタイプのフィールドをもつstructを宣言
	type Contact struct {
		email, address string
		phone          int
	}

	type Employee struct {
		name        string
		salary      int
		contactInfo Contact // embedded struct
	}

	// 2.embedded structに値設定する方法
	john := Employee{
		name:   "John Keller",
		salary: 4000,
		contactInfo: Contact{
			email:   "jkeller@company.com",
			address: "Street 20, London",
			phone:   0031545, // コンマ必須
		}, // コンマ必須
	}
	fmt.Printf("%#v\n", john) // main.Employee{name:"John Keller", salary:4000, contactInfo:main.Contact{email:"jkeller@company.com", address:"Street 20, London", phone:13157}}

	// 3. embedded structのフィールド値の取得
	fmt.Printf("Employee's email:%s\n", john.contactInfo.email) // Employee's email:jkeller@company.com

	// 4. embedded structの更新
	john.contactInfo.email = "new_email@company.com"
	fmt.Printf("Employee's email:%s\n", john.contactInfo.email) // Employee's email:new_email@company.com

	// 5.一度作ったembedded structを他の変数にも使う(john-> andrei)
	myContact := Contact{email: "andrei@domain.com", phone: 32123, address: "Bucharest Romania"}
	fmt.Println(myContact) // {andrei@domain.com Bucharest Romania 32123}
}
