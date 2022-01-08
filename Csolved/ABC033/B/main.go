package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	s, p := make([]string, n), make([]int, n)
	sp := 0
	for i := 0; i < n; i++ {
		s[i], p[i] = nextString(), nextInt()
		sp += p[i]
	}
	idx := -1
	for i := 0; i < n; i++ {
		if p[i] > sp/2 {
			idx = i
		}
	}
	if idx < 0 {
		fmt.Println("atcoder")
	} else {
		fmt.Println(s[idx])
	}
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
