package main

import "fmt"

// 0 1 1 2
// a b c
func main() {
	a := 0
	b := 1
	c := 1
	num := 125
	fmt.Printf("series:%d %d ", a, b)

	for true {
		c = b
		b = a + b
		if b >= num {
			break
		}

		a = c

		fmt.Printf("%d ", b)
	}
	fmt.Println()

}

func fibo(n int) {
	a := 0
	b := 1
	c := 1

	for true {
		c = b
		b = a + b
		if b >= n {

			break
		}
		a = c
		fmt.Printf("%d", b)
	}
}
