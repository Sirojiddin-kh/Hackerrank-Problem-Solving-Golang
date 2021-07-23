package main 

import (

	"fmt"
)
///
func unique(intSlice []int32) []int32 {
    keys := make(map[int32]bool)
    list := []int32{}	
    for _, entry := range intSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }    
    return list
}

func weightedUniformStrings(s string, queries []int32) []string {
	var (
		dict map[rune]int
		answer []int32
		result []string
		exis bool
	)

    for _, value := range s {
    	dict[value-96] += 1
    } 

    for key, value := range dict {

    	for i := 1; i <= value; i++ {
    		answer = append(answer, int32(key) * int32(i))
    	}
    }

    answer = unique(answer)


  	for _, value := range queries {
  		for _, value2 := range answer {

  			if value == value2 {
  				exis = true
  			} 
  		}

  		if exis {
  				result = append(result, "Yes") 
  		} else {
  			result = append(result, "No")
  		}
  		exis = false
  	}
  
    return result
}

func main() {

	s := "aaabbbbcccddd"
	queries := []int32{9, 7, 8, 12, 5}
	fmt.Println(weightedUniformStrings(s, queries))

}