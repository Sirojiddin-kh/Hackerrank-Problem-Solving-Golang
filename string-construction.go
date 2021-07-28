package main

import (
	"fmt"
)

func stringConstruction(s string) int32 {
	dict := make(map[string]uint32)

	for _, value := range s {

		dict[string(value)] += 1
	}

	return int32(len(dict))

}

func main() {

	s := "abcabc"
	fmt.Println(stringConstruction(s))
}
