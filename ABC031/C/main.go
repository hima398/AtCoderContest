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

	const INF = 251 // 2<=n<=50, -50 <= a[i] <= 50なので250より大きければ十分
	n := nextInt()
	a := nextIntSlice(n)
	ans := -INF
	for i := 0; i < n; i++ {
		mta, mao := -INF, -INF
		for j := 0; j < n; j++ {
			if j == i {
				continue
			}
			ta, ao := 0, 0
			pa := a[Min(i, j) : Max(i, j)+1]
			//fmt.Println(i, j, pa)
			for k := range pa {
				if k%2 == 0 {
					ta += pa[k]
				} else {
					ao += pa[k]
				}
			}
			if mao < ao {
				mta = ta //Max(mta, ta)
				mao = Max(mao, ao)
			}
		}
		ans = Max(ans, mta)
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
