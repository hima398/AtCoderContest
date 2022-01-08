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
	fact := make([]int, n+1)
	fact[0] = 1
	inv := make([]int, n+1)
	inv[0] = 1

	Comb := func(x, y int) int {
		if x < y {
			return 0
		} else {
			ret := fact[x] * inv[x-y]
			ret %= p
			ret *= inv[y]
			ret %= p
			return ret
		}
	}
	for i := 1; i <= n; i++ {
		fact[i] = fact[i-1] * i % p
		inv[i] = Inv(fact[i], p)
	}

	sort.Ints(a)

	ans := 0
	for i := n - 1; i >= 0; i-- {
		max := a[i] * Comb(i, k-1)
		ans += max
		ans %= p

		min := a[i] * Comb(n-i-1, k-1)
		ans -= min
		ans %= p
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

func Pow(x, y, p int) int {
	ret := 1
	for y > 0 {
		if y%2 == 1 {
			ret = ret * x % p
		}
		y >>= 1
		x = x * x % p
	}
	return ret
}

func Inv(x, p int) int {
	return Pow(x, p-2, p)
}
