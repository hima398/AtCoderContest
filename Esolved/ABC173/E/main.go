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

	const p = int(1e9) + 7

	n, k := nextInt(), nextInt()
	a := nextIntSlice(n)
	var plus, minus []int
	for _, ai := range a {
		if ai >= 0 {
			plus = append(plus, ai)
		} else {
			minus = append(minus, ai)
		}
	}
	np, nm := len(plus), len(minus)
	var solveLargeAbs bool
	if np > 0 {
		if n == k {
			solveLargeAbs = nm%2 == 0
		} else {
			solveLargeAbs = true
		}
	} else {
		// a[i]が全てマイナス
		solveLargeAbs = k%2 == 0
	}

	ans := 1
	if !solveLargeAbs {
		// 絶対値が小さい順にk個かける
		sort.Slice(a, func(i, j int) bool {
			return Abs(a[i]) < Abs(a[j])
		})
		for i := 0; i < k; i++ {
			ans *= a[i]
			ans %= p
		}
	} else {
		sort.Slice(plus, func(i, j int) bool {
			return Abs(plus[i]) > Abs(plus[j])
		})
		sort.Slice(minus, func(i, j int) bool {
			return Abs(minus[i]) > Abs(minus[j])
		})
		if k%2 == 1 {
			ans *= plus[0]
			ans %= p
			plus = plus[1:]
		}
		var s []int
		for len(plus) >= 2 {
			x := plus[0] * plus[1]
			plus = plus[2:]
			s = append(s, x)
		}
		for len(minus) >= 2 {
			x := minus[0] * minus[1]
			minus = minus[2:]
			s = append(s, x)
		}
		sort.Slice(s, func(i, j int) bool {
			return s[i] > s[j]
		})
		for i := 0; i < k/2; i++ {
			smp := s[i] % p
			ans *= smp
			ans %= p
		}
	}
	if ans < 0 {
		ans += p
	}
	ans %= p
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
