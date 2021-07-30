package main

import (
	"fmt"
)

func balancedSums(arr []int32) string {

	rightSum, leftSum := SumArr(arr), int32(0)

	for i := 0; i < len(arr); i++ {

		rightSum -= arr[i]

		if leftSum < rightSum {

			leftSum += arr[i]

		} else if leftSum == rightSum {

			return "YES"
		}
	}

	return "NO"

}

func SumArr(arr []int32) int32 {

	var sum int32

	for _, value := range arr {

		sum += value
	}

	return sum
}
