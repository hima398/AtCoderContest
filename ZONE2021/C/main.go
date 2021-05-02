package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

// コンテスト中DPで考えてみたので、そのアプローチで解いてみる
func FirstSolve() {
	n := nextInt()
	const k = 3
	const m = 5
	a := make([][]int, m)
	for i := 0; i < m; i++ {
		a[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			a[j][i] = nextInt()
		}
	}

	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = make([]int, 1<<m)
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j <= k; j++ {
			for k := 0; k < 1<<m; k++ {

			}
		}
	}

}
func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	const m = 5 //能力の種類
	a := make([][]int, m)
	for i := 0; i < m; i++ {
		a[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			a[j][i] = nextInt()
		}
	}
	check := func(x int) bool {
		b := make(map[int]bool)
		for j := 0; j < n; j++ {
			bit := 0
			for i := 0; i < m; i++ {
				bit = bit << 1
				if a[i][j] >= x {
					bit |= 1
				}
			}
			b[bit] = true
		}
		for k1 := range b {
			for k2 := range b {
				for k3 := range b {
					if k1|k2|k3 == 31 {
						return true
					}
				}
			}
		}
		return false
	}
	ok := 0
	ng := int(1e9) + 1
	for ng-ok > 1 {
		c := (ok + ng) / 2
		if check(c) {
			ok = c
		} else {
			ng = c
		}
	}
	fmt.Println(ok)
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
