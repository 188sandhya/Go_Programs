package main

import "fmt"

func main() {

	arr := []int{1, 3, 5, 3, 1, 4} // 4,5

	fmt.Println(firstUniqNum(arr))
	fmt.Println(isPalindrome("1221"))
}

func isPalindrome(str string) bool {
	for i := 0; i < len(str); i++ {
		j := len(str) - 1 - i
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

func firstUniqNum(s []int) int {
	rec := make(map[int]int)
	for _, b := range s {
		rec[b] = rec[b] + 1
	}

	for _, b := range s {
		if rec[b] == 1 {
			return b
		}
	}

	return -1
}
