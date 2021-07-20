package main

import (
	"fmt"
)

func main() {

	s := "hackerrank"
	c := "hhaacckkekraraannk"

	s1 := map[string]int 
	c1 := map[string]int

	for _, value := range s {
		s1[string(s)] += 1
	}
	for _, value := range c {
		c1[string(s)] + 1
	}
	fmt.Println(s1)
	fmt.Println(c1)
}