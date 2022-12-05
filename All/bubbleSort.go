package main

import "fmt"

func main() {
	arr := []int{22, 112, 67, 45, 2, 89}

	swapped := true

	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {

			if arr[i] < arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
				continue
			}

		}
	}

	fmt.Println("sorted:", arr)
}
