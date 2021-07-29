package main 

import (
	"fmt"
)

func icecreamParlor(m int32, arr []int32) []int32 {
    
    var result []int32 

    for i := 0; i < len(arr); i++ {

    	for j := i + 1; j < len(arr); j++ {

    		if arr[i] + arr[j] == m {

    			result = append(result, int32(i+1), int32(j+1))
    		}
    	}
    }
    return result

}

func main() {

	a := []int32{1, 4, 5, 3, 2}
	m := int32(4)
	fmt.Println(icecreamParlor(m, a))
}