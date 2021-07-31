package main

import (
	"fmt"
	"math"
	"sort"
)

func marcsCakewalk(calorie []int32) int64 {

	var minMiles int64

	arr := convertInt(calorie)
	fmt.Println("%T", arr)

	for i := 0; i < len(calorie); i++ {

		minMiles += int64(math.Pow(float64(2), float64(i))) * int64(arr[i])

	}
	return minMiles
}
func convertInt(array []int32) []int {

	var result []int

	for _, value := range array {

		result = append(result, int(value))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(result)))

	return result
}

func main() {

	arr := []int32{5, 10, 7}

	fmt.Println(marcsCakewalk(arr))
}
