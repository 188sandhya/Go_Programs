package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(timeConversion("08:05:45PM"))

}
func timeConversion(s string) string {
	// Write your code here

	layout1 := "03:04:05PM"
	layout2 := "15:04:05"
	t, err := time.Parse(layout1, s)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(t.Format(layout1))
	// fmt.Println(t.Format(layout2))
	return t.Format(layout2)

}
