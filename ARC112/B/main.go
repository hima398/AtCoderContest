package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	a := make([]int, n)
	m := make(map[int]int)
	max := 0
	for i := 0; i < n; i++ {
		a[i] = nextInt()
		m[a[i]]++
		max = Max(max, a[i])
	}
	sort.Ints(a)
	premin := k
	ans := 0
	for i := 0; i <= max; i++ {
		min := Min(premin, m[i])
		ans += (premin - min) * i
		premin = min
	}
	if premin > 0 {
		ans += (max + 1) * premin
	}
	fmt.Println(ans)
}
