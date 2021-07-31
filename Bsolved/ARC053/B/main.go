package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var sc = bufio.NewScanner(os.Stdin)

func FirstSolve(s string) int {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	var q []int
	for _, v := range m {
		q = append(q, v)
	}
	sort.Ints(q)
	//fmt.Println(q)
	var answers []int
	for len(q) > 0 {
		// dequeue
		v := q[0]
		q = q[1:]
		if v%2 == 0 {
			if len(answers) == 0 {
				answers = append(answers, v)
			} else {
				answers[0] += v
			}
		} else {
			answers = append(answers, v)
		}
		sort.Ints(answers)
	}
	//fmt.Println(answers)
	return answers[0]
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextString()
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	no := 0
	for _, v := range m {
		if v%2 == 1 {
			no++
		}
	}
	var ans int
	if no == 0 {
		ans = len(s)
	} else {
		ans = 2*Floor(len(s)-no, (2*no)) + 1
	}
	fmt.Println(ans)
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Floor(x, y int) int {
	return x / y
}
