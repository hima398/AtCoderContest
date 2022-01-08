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

	const INF = int(1e9)
	n := nextInt()
	a := nextIntSlice(2 * n)
	memo := make([][]int, 2*n)
	v := make([][]bool, 2*n)
	for i := 0; i < 2*n; i++ {
		memo[i] = make([]int, 2*n)
		v[i] = make([]bool, 2*n)
		for j := 0; j < 2*n; j++ {
			memo[i][j] = INF
		}
	}
	var F func(l, r int) int
	F = func(l, r int) int {
		//fmt.Println(l, r)
		if l == r {
			return 0
		}
		if v[l][r] {
			return memo[l][r]
		}
		if r-l == 1 {
			memo[l][r] = Abs(a[l] - a[r])
			v[l][r] = true
			return memo[l][r]
		}
		ret := INF
		for i := l + 1; i < r-1; i += 2 {
			ret = Min(ret, F(l, i)+F(i+1, r))
		}
		//l-rの範囲のうち真ん中を消して両端を消すパターン
		if (r-l+1)%2 == 0 {
			ret = Min(ret, F(l+1, r-1)+Abs(a[l]-a[r]))
		}
		memo[l][r] = ret
		v[l][r] = true
		return memo[l][r]
	}
	ans := F(0, 2*n-1)
	//fmt.Println(memo)
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

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
