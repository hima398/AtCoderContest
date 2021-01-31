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

	n, m, k := nextInt(), nextInt(), nextInt()
	sa := make([]int, n+1)
	for i := 0; i < n; i++ {
		a := nextInt()
		sa[i+1] += sa[i] + a
	}
	sb := make([]int, m+1)
	for i := 0; i < m; i++ {
		b := nextInt()
		sb[i+1] += sb[i] + b
	}
	ans := 0
	j := m
	for i := 0; i <= n; i++ {
		if sa[i] > k {
			break
		}
		kb := k - sa[i]
		for sb[j] > kb {
			j--
		}
		ans = Max(ans, i+j)
	}
	fmt.Println(ans)
}
