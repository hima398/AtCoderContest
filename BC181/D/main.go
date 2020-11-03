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

func Exp(m, n int) int {
	for i := 0; i < n; i++ {
		m *= n
	}
	return m
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
	// string -> []int
	// ex. "1234" -> []int{1, 2, 3, 4}
	si := []int{}
	for i := 0; i < len(s); i++ {
		v, _ := strconv.Atoi(string(s[i]))
		si = append(si, v)
	}
	return ""
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(ProblemD(s))
}
