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

func CreateRoute(s string) []string {
	var ss []string
	if len(s) == 1 {
		return []string{s + "1"}
	}
	for i := 0; i < len(s); i++ {
		sss := CreateRoute(DeleteS(s, i))
		for j := 0; j < len(sss); j++ {
			sss[j] = string(s[i]) + sss[j]
		}
		ss = append(ss, sss...)
	}
	return ss
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	t := make([][]int, n)
	baseRoute := ""
	for i := 0; i < n; i++ {
		t[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&t[i][j])
		}
		baseRoute += strconv.Itoa(i + 1)
	}
	route := CreateRoute(DeleteS(baseRoute, 0))
	for i := 0; i < len(route); i++ {
		route[i] = "1" + route[i]
	}

	var result int
	for i := 0; i < len(route); i++ {
		d := 0
		for j := 0; j < len(route[i])-1; j++ {
			current, _ := strconv.Atoi(string(route[i][j]))
			next, _ := strconv.Atoi(string(route[i][j+1]))
			d += t[current-1][next-1]
		}
		if d == k {
			result++
		}
	}
	fmt.Println(result)
}
