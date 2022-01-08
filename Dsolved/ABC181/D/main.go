package main

import (
	"fmt"
	"strconv"
)

func Swap(s string) string {
	r := []rune(s)
	r[0], r[1] = r[1], r[0]
	return string(r)
}

func ProblemD(s string) bool {

	if len(s) == 1 {
		return s == "8"
	}
	if len(s) == 2 {
		n1, _ := strconv.Atoi(s)
		n2, _ := strconv.Atoi(Swap(s))
		return n1%8 == 0 || n2%8 == 0
	}
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	for i := 112; i < 1000; i += 8 {
		m2 := map[byte]int{}
		for k, v := range m {
			m2[k] = v
		}
		s2 := strconv.Itoa(i)
		for j, _ := range s2 {
			m2[s2[j]]--
		}
		b := true
		for _, v := range m2 {
			b = b && v >= 0
		}
		if b {
			return b
		}
	}
	return false
}

func main() {
	var s string
	fmt.Scan(&s)
	if ProblemD(s) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
