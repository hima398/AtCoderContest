package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

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

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	s := nextString()
	m := make(map[string]int)
	for i := 0; i < n; i++ {
		m[string(s[i])]++
	}
	//fmt.Println(m)
	nw := 0
	for i := 0; i < m["R"]; i++ {
		if s[i] == 'W' {
			nw++
		}
	}
	fmt.Println(Min(m["W"], nw))
}
