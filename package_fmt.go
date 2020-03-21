package main

import (
	"fmt"
)

func main() {
	a, b := 4, 6
	fmt.Println("Sum:", a+b, "Mean Value:", (a+b)/2)

	// 1.Printf フォーマット
	fmt.Printf("Your age is %d\n", 21) // Printfの改行には/nが必要
	fmt.Printf("x is %d, y is %f\n", 5, 6.8)

	// 2.エスケープsequence
	fmt.Printf("He says: \"Hello Go!\"\n")

	// 3. %d, %f, %s(文字), %q(引用)
	figure := "Circle"
	radius := 5
	pi := 3.14159

	_, _ = figure, pi

	fmt.Printf("Radius is %d\n", radius)
	fmt.Printf("Radius is %+d\n", radius)

	fmt.Printf("Pi constant is %f\n", pi)

	fmt.Printf("The diameter of a %s with a Radius of %d is %f\n", figure, radius, float64(radius)*2*pi)

	fmt.Printf("This is a %q\n", figure) // Circle → "Circle" 引用句

	// 4. %v -> どのtypeにも変換可能
	fmt.Printf("The diameter of a %v with a Radius of %v is %v\n", figure, radius, float64(radius)*2*pi)

	// 5. %T : タイプを表示
	fmt.Printf("figure is a type of %T\n", figure)
	fmt.Printf("radius is a type of %T\n", radius)
	fmt.Printf("pi is a type of %T\n", pi)

	// 6. %t : boolean
	closed := true
	fmt.Printf("File closed: %t\n", closed)

	// 7.%b : base 2 (binary) ,  %08b : 8桁, %x : hexidecimal(base 16)
	fmt.Printf("%b \n", 55)   // 32+16+4+2+1
	fmt.Printf("%08b \n", 55) //
	fmt.Printf("%x \n", 100)  // 16*6 + 4

	// 8.小数点(float64)
	x := 3.4
	y := 6.9

	fmt.Printf("x * y = %f\n", x*y)
	fmt.Printf("x * y = %.3f\n", x*y) // 小数点以下3桁
}
