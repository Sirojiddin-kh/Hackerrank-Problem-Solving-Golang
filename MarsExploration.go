package main 

import (

	"fmt"
)

func marsExploration(s string) int32 {
    var delta int32
    for j := range s {
    if s[j] != "SOS"[j%3] {
    	fmt.Println("SOS"[j%3])
        delta++
    }
    }
    return delta

}



