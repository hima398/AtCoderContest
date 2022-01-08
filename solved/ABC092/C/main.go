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
	a := nextIntSlice(n)
	a = append([]int{0}, a...)
	a = append(a, 0)
	s := Abs(a[1]) + Abs(a[n])
	for i := 1; i < n; i++ {
		s += Abs(a[i+1] - a[i])
	}
	ans := make([]int, n)
	for i := 1; i <= n; i++ {
		ans[i-1] = s
		ans[i-1] -= Abs(a[i-1] - a[i])
		ans[i-1] -= Abs(a[i] - a[i+1])
		ans[i-1] += Abs(a[i-1] - a[i+1])
	}
	// Abs(0 - (-1)) + Abs((-1) - 3) + Abs(3 - 5) + Abs(5 - 0) = 1 + 4 + 2 + 5 = 12
	// 5 - 3 + 3 - (-1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := 0; i < n; i++ {
		fmt.Fprintln(out, ans[i])
	}
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
