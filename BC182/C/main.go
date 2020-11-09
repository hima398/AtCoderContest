package main

import (
	"fmt"
	"strconv"
)

func DeleteS(s string, idx int) string {
	var result string
	result = s[:idx]
	result += s[idx+1:]
	return result
}

func To3(n, d int) int {
	if n > 0 && n <= 9 {
		if n%3 == 0 {
			return d
		} else {
			return -1
		}
	} else {
		s := strconv.Itoa(n)
		detected := -1
		for i := 0; i < len(s); i++ {
			subS := DeleteS(s, i)
			subN, _ := strconv.Atoi(subS)
			if subN%3 == 0 {
				return d
			} else {
				t := To3(subN, d+1)
				if t == -1 {
					continue
				} else {
					detected = t
				}
			}
		}
		return detected
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	// そのままわれるかをチェックする
	if n%3 == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(To3(n, 1))
	}
}
