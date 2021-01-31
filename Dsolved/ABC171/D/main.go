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

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	ai := make([]int, n)
	m := make(map[int]int)
	s := 0
	for i := 0; i < n; i++ {
		a := nextInt()
		ai[i] = a
		m[a]++
		s += a
	}
	ans := s
	q := nextInt()
	for i := 0; i < q; i++ {
		b, c := nextInt(), nextInt()
		t := m[b]
		m[b] = 0
		m[c] += t
		ans -= b * t
		ans += c * t
		fmt.Println(ans)
	}
}
