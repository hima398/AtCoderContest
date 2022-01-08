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

	n, k := nextInt(), nextInt()
	v := nextIntSlice(n)
	ans := 0
	for l := 0; l <= n; l++ {
		for r := 0; r <= n; r++ {
			if l+r > n {
				continue
			}
			rem := k - (l + r)
			if rem < 0 {
				continue
			}
			var jwl []int
			for i := 0; i < l; i++ {
				jwl = append(jwl, v[i])
			}
			for i := 0; i < r; i++ {
				jwl = append(jwl, v[n-i-1])
			}
			sort.Ints(jwl)
			//fmt.Println(jwl, rem)
			//jwl = jwl[rem:]
			s := 0
			for i := range jwl {
				s += jwl[i]
			}
			for i := 0; i < Min(rem, len(jwl)); i++ {
				if jwl[i] > 0 {
					break
				}
				s -= jwl[i]
			}
			ans = Max(ans, s)
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
