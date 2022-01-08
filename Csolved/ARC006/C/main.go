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
	w := nextIntSlice(n)
	q := []int{w[0]}
	for i := 1; i < n; i++ {
		bn := len(q)
		idx := sort.Search(bn, func(j int) bool {
			return w[i] <= q[j]
		})
		if idx >= 0 && idx < bn {
			q[idx] = w[i]
		} else {
			q = append(q, w[i])
		}
		sort.Ints(q)
	}
	fmt.Println(len(q))
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextIntSlice(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	return s
}
