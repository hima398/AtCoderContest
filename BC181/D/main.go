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

func PermuteS(s string) []string {
	if len(s) == 1 {
		return []string{s}
	}
	var result []string
	for i, _ := range s {
		ss := PermuteS(DeleteS(s, i))
		for _, sss := range ss {
			result = append(result, string(s[i])+sss)
		}
	}
	return result
}

func ProblemD(s string) string {

	permute := PermuteS(s)
	for _, v := range permute {
		n, _ := strconv.Atoi(v)
		if n%8 == 0 {
			return "Yes"
		}
	}
	return "No"
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(ProblemD(s))
}
