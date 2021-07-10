package main

import (
	"fmt"
	"math"
)

func main() {
	var n int32 = 5
	var c []int32 = []int32{0, 4}
	var perDistance [][]int32
	var nearestDistance []int32
	var distance []int32

	for i := 0; int32(i) < (n - 1); i++ {
		for j := 0; j < len(c); j++ {
			distance = append(distance, int32(math.Abs(float64(int32(i)-c[j]))))
		}
		perDistance = append(perDistance, distance)
		distance = []int32{}

	}
	fmt.Println(perDistance)
	var min int32
	for i := 0; i < len(perDistance); i++ {
		min = perDistance[i][0]
		//fmt.Println(min)
		for j := 0; j < len(perDistance[i]); j++ {

			if perDistance[i][j] > min {
				min = perDistance[i][j]
				fmt.Println(min)
				nearestDistance = append(nearestDistance, min)
			}

		}
	}
	fmt.Println(nearestDistance)
}
