package main

import (

	"fmt"
	"math"
)

func minimumAbsoluteDifference(arr []int32) int32 {
    
	var differences []int32 

	for i := 0; i < len(arr) - 1; i++ {

		for j := i + 1; j < len(arr); j++ {

			differences = append(differences, int32(math.Abs(float64(int32(arr[i]) - int32(arr[j]))))
		}

	}

	minimum := int32(differences[0])

	for _, value := range differences {

		if value > minimum {

			minimum = value
		}
	}

	return minimum

}

func main() {

	a := []int32{-2, 2, 4}

	fmt.Println(minimumAbsoluteDifference(a))
}