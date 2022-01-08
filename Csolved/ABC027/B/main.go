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
	sa := make([]int, n+1)
	for i := range a {
		sa[i+1] += sa[i] + a[i]
	}
	s := sa[n]
	if s%n != 0 {
		fmt.Println(-1)
		return
	}
	ans := 0
	for i := 1; i <= n; i++ {
		sl, sr := sa[i], sa[n]-sa[i]
		if sl*(n-i) != sr*i {
			ans++
		}
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
