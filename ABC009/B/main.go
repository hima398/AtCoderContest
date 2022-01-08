package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		ai := nextInt()
		m[ai]++
	}
	var s []int
	for k := range m {
		s = append(s, k)
	}
	sort.Ints(s)
	// n >= 2, 全ての料理の金額が同じでないことからlen(s) >= 2の想定
	ans := s[len(s)-2]
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
