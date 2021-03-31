package main

import "fmt"

func main() {

	Student := make(map[string]int)
	Student["sameer"] = 45

	fmt.Println(Student["sameer"])

	// map inside map

	superhero := map[string]map[string]string{
		"Superman": {
			"Realname": "Kent",
			"City":     "Metropolis",
		},
	}

	fmt.Println(superhero["Superman"]["City"])
}
