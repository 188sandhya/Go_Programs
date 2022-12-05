package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int64 = 1
	result := flippingBits(n)

	fmt.Println(result)

}

func flippingBits(n int64) int64 {
	nBin := fmt.Sprintf("%032b", n)
	fmt.Println(nBin)

	var sb strings.Builder

	for _, v := range nBin {
		if string(v) == "0" {
			sb.WriteString("1")
		} else if string(v) == "1" {
			sb.WriteString("0")
		}
	}

	newNumBin := sb.String()

	out, _ := strconv.ParseUint(newNumBin, 2, 32)

	return int64(out)
}
