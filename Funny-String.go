package main

import (

	"fmt"
	"math"
)

func funnyString(s string) string {

	reverse := ""
	var (
		origin, reversed, dif1, dif2 []int
		count int
	)

	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	for i := 0; i < len(s); i++ {
		origin = append(origin, int(s[i]))
		reversed = append(reversed, int(reverse[i]))
	}


	for i := 0; i < len(origin)- 1; i++ {
	
		dif1 = append(dif1, int(math.Abs(float64(origin[i]- origin[i+1]))))
		dif2 = append(dif2, int(math.Abs(float64(reversed[i]- reversed[i+1]))))

	}


	for i := 0; i < len(dif1); i++ {
		if dif1[i] == dif2[i] {
			count += 1
		}
	}

	if count == len(dif1) {
		return "Funny"
	} else {
		return "Not Funny"
	}

}

func main() {


	s := "ovyvzvptyvpvpxyztlrztsrztztqvrxtxuxq"
	fmt.Println(funnyString(s))

}