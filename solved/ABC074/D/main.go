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
	a := make([][]int, n)
	f := make([][]bool, n)
	for i := 0; i < n; i++ {
		a[i] = nextIntSlice(n)
		f[i] = make([]bool, n)
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i == j || i == k || j == k {
					continue
				}
				if a[i][j] > a[i][k]+a[k][j] {
					fmt.Println(-1)
					return
				}
				if a[i][j] == a[i][k]+a[k][j] {
					//ans -= a[i][j]
					f[i][j] = true
				}
			}
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if !f[i][j] {
				ans += a[i][j]
			}
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
