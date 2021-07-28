package main

import (
	"fmt"
)

func gameOfthrones(s string) string {

	dict := make(map[string]uint32)

	var pairs []uint32

	for _, value := range s {

		dict[string(value)] += 1

	}
	for _, value := range dict {

		if value % 2 == 0 {

			pairs = append(pairs, value)
		}
	}
	fmt.Println(dict)
	fmt.Println(pairs)

	if len(pairs) != 0 {

		return "YES"
	} else {

		return "NO"
	}
}

func main() {

	s := ""
	fmt.Println(gameOfthrones(s))
}
