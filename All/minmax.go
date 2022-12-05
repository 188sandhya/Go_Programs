package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int32{256741038, 623958417, 467905213, 714532089, 938071625}
	// arr := []int32{256, 61123, 41169, 71114, 938071625}
	miniMaxSum(arr)
}
func miniMaxSum(arr []int32) {
	// Write your code here
	sort.Slice(arr, func(i, j int) bool {
		return arr[j] > arr[i]
	})

	fmt.Println("sorted:", arr)
	lenArr := len(arr)
	minArr := arr[:lenArr-1]
	fmt.Println("minarr:", minArr)
	min := sum(minArr)
	maxArr := arr[1:]

	fmt.Println("maxarr:", maxArr)
	max := sum(maxArr)
	fmt.Println(min, max)
}

func sum(n []int32) int32 {
	var result int32 = 0

	for _, v := range n {
		result = result + v
		fmt.Println("result:", result)
	}

	return result
}
