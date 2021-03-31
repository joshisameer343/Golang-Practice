package main

import "fmt"

func main() {
	var a [5]int

	b := [5]int{1, 2, 3, 4, 5}

	a[0] = 2
	a[1] = 4

	fmt.Println(a[1])
	fmt.Println(b[1])

	for i, value := range b {

		fmt.Println(value, i)
	}

	for _, value := range b {

		fmt.Println(value)
	}
}
