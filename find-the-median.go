package main 

import (
	"fmt"
	"sort"
)
func sorting(arr []int32) []int32 {

	var arrInt []int
	var arrInt32 []int32

	for _, value := range arr {

		arrInt = append(arrInt, int(value))
	}
	sort.Ints(arrInt)

	for _, value := range arrInt {

		arrInt32 = append(arrInt32, int32(value))
	}
	return arrInt32

}

func findMedian(arr []int32) int32 {

	var ( 
		middle, result int32
		median []int32
	)

	arr = sorting(arr)

	middle = int32(len(arr)/2)

	median = arr[middle : middle + 1]

	for _, value := range median {

		result = value
	}



	return int32(result)

}

func main() {

	arr := []int32{0, 1, 2, 4, 6, 5, 3}
	arr = sorting(arr)

	fmt.Println(arr)

	fmt.Println(findMedian(arr))
}