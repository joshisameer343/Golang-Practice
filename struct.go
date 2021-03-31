package main

import (
	"fmt"
	"strconv"
)

func main() {

	sam1 := Info{34, "sameer"}
	fmt.Println(sam1.age)
	fmt.Println(sam1.name)
	fmt.Println(sam1.printInfo())

}

type Info struct {
	age  int
	name string
}

func (x *Info) printInfo() string {

	return ("name is " + x.name + " age is " + strconv.Itoa(x.age))
}
