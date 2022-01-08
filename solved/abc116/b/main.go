package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func F(n int) int {
	if n%2 == 0 {
		return n / 2
	} else {
		return 3*n + 1
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextInt()
	pre, next := s, 0
	m := make(map[int]bool)
	m[pre] = true
	i := 1
	for {
		i++
		next = F(pre)
		if m[next] {
			fmt.Println(i)
			return
		}
		m[next] = true
		pre = next
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
