package main

import (
	"fmt"
	"math"
)

func main() {
	var arr []int = []int{7, 1, 3, 4, 1, 7}
	var allDistance []int
    for key1, value1 := range arr {
        for key2, value2 := range arr {
            if value1 == value2 && key2 != key1{
                allDistance = append(allDistance, int(math.Abs(float64(key2-key1)))) 
            }
        }
    }
    var min int = allDistance[0]

    for _, value := range allDistance {

    	if value < min {
    		min = value
    	}
    }
    fmt.Println(min)
}