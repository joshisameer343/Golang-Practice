package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Create("sam.txt")

	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("This is some random text")

	file.Close()

	stream, err := ioutil.ReadFile("sam.txt")

	if err != nil {
		log.Fatal(err)
	}

	s1 := string(stream)

	fmt.Println(s1)

}
