package main

import (
	"fmt"
)

func anagram(s string) int32 {
	if len(s)%2 != 0 {
		return -1
	}
	s1 := s[len(s)/2:]
	s2 := s[:len(s)/2]
	
	m := make(map[rune]int)
	for _, c := range s1 {
		m[c]++
	}
	fmt.Println(m)
	for _, c := range s2 {
		m[c]--
	}
	t := int32(0)
	fmt.Println(m)
	for _, v := range m {
		if v < 0 {
			continue
		}
		t += int32(v)
	}
	return t
}

func main() {

	s := "xaxbbbxx"

	fmt.Println(anagram(s))
}
