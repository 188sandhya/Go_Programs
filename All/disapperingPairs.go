package main

import (
	"fmt"
)

func main() {
	str := "ACCABBCA"
	result := disappear((str))
	fmt.Println("final string:", result)

}

func disappear(str string) string {
	// count := make(map[string]int)

	// var res string
	// if !strings.Contains(str, "A") || !strings.Contains(str, "B") || !strings.Contains(str, "C") {
	// 	return str
	// }

	// for _, v := range str {
	// 	count[string(v)] = count[string(v)] + 1
	// }

	// for i, v := range count {
	// 	if v%2 == 0 {
	// 		delete(count, i)
	// 	} else {
	// 		res = res + i
	// 	}
	// }

	var res string

	for i := 0; i <= len(str)-2; i++ {
		if str[i] == str[i+1] {
			continue
		} else {
			res = res + string(str[i])
		}
	}

	return res

}
