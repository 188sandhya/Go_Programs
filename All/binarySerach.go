package main

import "fmt"

// to perform binary serrach array must be sorted
// iterative method

const (
	i = 5
	j
	k
)

func printType(a interface{}) {
	fmt.Printf("type:%T\n", a)
}

func main() {

	var i interface{}
	i = "heyy"
	i = 9.888
	printType(i)

	fmt.Println("checj i j k:", i, j, k)
	arr := []int{23, 34, 55, 67, 89, 90}
	x := 55
	low := 0
	high := len(arr) - 1
	ind := findNumber(arr, x, low, high)
	fmt.Println("found at :", ind)
}

func findNumber(arr []int, x, low, high int) int {

	for low <= high {

		mid := (low + high) / 2

		if arr[mid] == x {
			return mid
		} else if x < arr[mid] {
			high = mid - 1

		} else {
			low = mid + 1
			// return findNumb/er(arr, x, low, high)
		}

	}

	return -1
}
