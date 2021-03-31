package main

import "fmt"

func main() {

	var age int = 21

	if age > 18 {
		fmt.Println("eligible")
	} else {
		fmt.Println("not eligible")
	}

	switch age {

	case 18:
		fmt.Println("not eligible")

	case 19:
		fmt.Println(" eligible")

	case 21:
		fmt.Println(" eligible")

	default:
		fmt.Println("not eligible")

	}
}
