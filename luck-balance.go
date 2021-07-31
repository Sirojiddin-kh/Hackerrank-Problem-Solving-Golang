package main

import (

	"fmt"
	"sort"
)

func luckBalance(k int32, contests [][]int32) int32 {
   sort.SliceStable(contests, func(i, j int) bool { return contests[i][0] > contests[j][0] })

    losts := int32(0)
    luck := int32(0)

    for _, v := range contests {
        if losts+v[1] > k {
            luck -= v[0]
        } else {
            losts += v[1]
            luck += v[0]
        }
    }
    return luck

}

func main() {

	constest := [][]int32{{5, 1}, {2, 1}, {1, 1}, {8, 1}, {10, 0}, {5, 0}}

	fmt.Println(luckBalance(3, constest))
}