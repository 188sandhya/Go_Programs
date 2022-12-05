package main

import "fmt"

func main() {

	arr := []int32{1, 2, 1, 9, 2, 1}
	fmt.Println(lonelyinteger(arr))

}

func lonelyinteger(a []int32) int32 {
	// Write your code here
	aMap := make(map[int32]int32)

	for _, v := range a {
		fmt.Println("aMap[v]:", aMap[v])
		aMap[v] = aMap[v] + 1
	}

	fmt.Println("amap:", aMap)

	for k, v := range aMap {
		if v == 1 {
			return k
		}
	}

	return 0
}
