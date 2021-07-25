package main

import (

	"fmt"
	"math"
)
func reversed(s string) string {

	var reverse string

	for i := len(s) - 1; i >=0 ; i-- {
		reverse += string(s[i])
	}
	return reverse
}

func theLoveLetterMystery(s string) int32 {

    var count int32

	if s == reversed(s) {

		return 0

	} else {
		j := int32(0)
		for i := len(s) -1; int32(i) >= int32(len(s) - len(s) / 2); i-- {

			count += int32(math.Abs(float64(int32(s[j]) - int32(s[i]))))
			j++
		}

	}

	return count
}

func main() {

	a := "abac"
	fmt.Println(theLoveLetterMystery(a))
}
