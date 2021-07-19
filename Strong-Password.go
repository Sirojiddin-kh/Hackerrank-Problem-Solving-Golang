package main

import (
	"fmt"
	"unicode"
	"strings"
	"math"
)

func minimumNumber(n int32, password string) int32 {

	required := map[string]bool {
		"number" : false,
		"lowerCase": false,
		"upperCase": false,
		"specialChar": false,
		"length": false,
	}

	length := len(password)

	for _, value := range password {

		if unicode.IsLower(value) {
			required["lowerCase"] = true
		}
		if unicode.IsUpper(value) {
			required["upperCase"] = true
		}
		if unicode.IsNumber(value) {
			required["number"] = true
		}
		if strings.ContainsRune("!@#$%^&*()-+", value) {
			required["specialChar"] = true
		}
	}
	count := int32(4)
	for _, value := range required {
		if value == true {
			count -= 1
		}
	}
	return int32(math.Max(float64(count), float64(6-length)))
}

func main() {
	a := int32(11)
	password := "#HackerRank"
	answer := minimumNumber(a, password)
	fmt.Println(answer)
}