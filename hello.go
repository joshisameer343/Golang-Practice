package main

import "fmt"

func main() {

	// declaring variable
	var a int = 5
	var b float32 = 4.32
	var b1 = 333
	const pi float64 = 3.14
	p := 23
	isbool := true
	isbool1 := false

	// declaring two variables
	x, y := 14, 15

	var (
		c = 8
		d = 7
	)

	fmt.Println(a, b, pi, c, d, x, y, p)

	// printing type
	fmt.Printf("%T", b1)

	// operator
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)

	fmt.Println(isbool && isbool1)
	fmt.Println(isbool || isbool1)
	fmt.Println(!isbool)

	// Pointers

	fmt.Println(&x)

}
