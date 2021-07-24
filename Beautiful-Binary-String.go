package main

import (
	"fmt"
)

func beautifulBinaryString(b string, n int32) int32 {
    
    var (
        substring string = ""
        count int32
    )

    for i := 0; int32(i) < n - 2; i++ {

        substring += string(b[i]) + string(b[i+1]) + string(b[i+2])

        if substring == "010" {
            count += 1 
            i = i + 2
        }
        substring = ""
    }
    return count
}
