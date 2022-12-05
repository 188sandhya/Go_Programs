package main

import (
	"fmt"
	"strings"
)

func main() {

	occurred := map[int]bool{}
	result := []int{}
	arr := []int{2, 4, 2, 6, 8, 4, 5, 2}
	for e := range arr {

		// check if already the mapped
		// variable is set to true or not
		if occurred[arr[e]] != true {
			occurred[arr[e]] = true

			// Append to result slice.
			result = append(result, arr[e])
		}
	}
	fmt.Println("uniuq ints:", result)
	str := "AABBCCDDD"
	fmt.Println("unique chars:", removeDuplicates(str))

}

func removeDuplicates(s string) string {
	res := make(map[rune]bool)
	var result []string
	for _, c := range s {
		if value := res[c]; !value {
			res[c] = true
			result = append(result, string(c))
		}
	}

	return strings.Join(result, "")
}
