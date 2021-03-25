package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	s := nextString()
	p := make([]int, n+1)
	min := len(s) + 1
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			p[i+1] = p[i] + 1
		} else {
			// s[i] == ')'
			p[i+1] = p[i] - 1
		}
		min = Min(min, p[i+1])
	}
	//fmt.Println(-min, p[len(s)]-min)
	no := -Min(min, 0)
	nc := p[len(s)] + no
	open := strings.Repeat("(", no)
	close := strings.Repeat(")", nc)
	fmt.Println(open + s + close)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
