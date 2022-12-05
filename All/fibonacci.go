package main

import "fmt"

func main() {
	fibo(20)
}

// a b c
// 0 1 1
func fibo(n int) {
	a := 0
	b := 1
	c := 0

	fmt.Printf("fibo series:%d %d", a, b)
	for true {
		c = b
		b = a + b
		if b >= n {
			fmt.Println()
			break
		}

		a = c
		fmt.Printf(" %d", b)

	}
}
