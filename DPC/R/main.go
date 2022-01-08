package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveHonestly(n, k int, a [][]int) int {

	const p = int(1e9) + 7
	e1, e2 := make([][]int, n), make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] > 0 {
				e1[i] = append(e1[i], j)
				e2[j] = append(e2[j], i)
			}
		}
	}
	dp := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for kk := 1; kk <= k; kk++ {
		for i := 0; i < n; i++ {
			for _, j := range e2[i] {
				dp[kk][i] += dp[kk-1][j]
				dp[kk][i] %= p
			}
		}
		fmt.Println(dp[kk])
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans += dp[k][i]
		ans %= p
	}
	return ans
}

func MulMatrix(n int, a, b [][]int, p int) [][]int {
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				ret[i][j] += a[i][k] * b[k][j]
				ret[i][j] %= p
			}
		}
	}
	return ret
}

func PowMatrix(n int, a [][]int, x, p int) [][]int {
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		ret[i][i] = 1
	}
	for x > 0 {
		if x%2 == 1 {
			ret = MulMatrix(n, ret, a, p)
		}
		x >>= 1
		a = MulMatrix(n, a, a, p)
	}
	return ret
}

func Solve(n, k int, a [][]int) int {
	const p = int(1e9) + 7

	dpn := PowMatrix(n, a, k, p)

	var ans int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ans += dpn[i][j]
			ans %= p
		}
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextIntSlice(n)
	}
	//ans := SolveHonestly(n, k, a)
	ans := Solve(n, k, a)
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
