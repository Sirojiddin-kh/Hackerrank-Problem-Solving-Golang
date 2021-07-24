package main

import (

	"fmt"
	"math"
)

func alternatingCharacters(s string) int {
    
    var (
    	deletions int
    	sequences []int
    )

    for i := 0; i < len(s)-1; i++ {

    	sequences = append(sequences, int(math.Abs(float64(s[i+1] - s[i]))))
    }
    
    for _, value := range sequences {

    	if value == 0 {
    		deletions +=1
    	}
    }
    return deletions
}

func main() {

	s := "AAABBB"
	fmt.Println(alternatingCharacters(s))
}