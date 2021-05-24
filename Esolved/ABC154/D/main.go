package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	a := nextIntSlice(n)
	sa := make([]float64, n+1)
	for i := 0; i < n; i++ {
		sa[i+1] = float64((a[i] + 1) * a[i])
		sa[i+1] /= 2.0 * float64(a[i])
	}
	for i := 0; i < n; i++ {
		sa[i+1] += sa[i]
	}
	ans := 0.0
	for i := k; i <= n; i++ {
		e := sa[i] - sa[i-k]
		ans = math.Max(ans, e)
	}
	fmt.Println(ans)
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
