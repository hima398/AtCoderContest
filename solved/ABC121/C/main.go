package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Pair struct {
	x, y int
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)
	n, m := nextInt(), nextInt()
	p := make([]Pair, n)
	for i := 0; i < n; i++ {
		p[i].x, p[i].y = nextInt(), nextInt()
	}
	sort.Slice(p, func(i, j int) bool { return p[i].x < p[j].x })
	num := 0
	ans := 0
	for i := 0; i < n; i++ {
		min := Min(m-num, p[i].y)
		num += min
		ans += p[i].x * min
		if num >= m {
			break
		}
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
