package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveHonestly(n int, a [][]int) int {
	const p = int(1e9) + 7
	//男性i(0-indexed)のペアを決めた時、女性の中でペアが決まったかどうかがbitで管理された状況になっている組み合わせ数
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 1<<n)
	}
	dp[0][0] = 1
	// O(N**2 * 2**N)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] == 1 {
				for k := 0; k < 1<<n; k++ {
					//j番目の女性のペアがまだ決まっていない
					if (k>>j)&1 == 0 {
						nj := k | (1 << j)
						dp[i+1][nj] += dp[i][k]
						dp[i+1][nj] %= p
					}
				}
			}
		}
		//fmt.Println(dp[i+1])
	}
	return dp[n][1<<n-1]
}

func Solve(n int, a [][]int) int {
	const p = int(1e9) + 7

	dp := make([]int, 1<<n)
	dp[0] = 1
	for pat := 0; pat < 1<<n; pat++ {
		i := bits.OnesCount(uint(pat))
		for j := 0; j < n; j++ {
			if pat>>j&1 == 0 && a[i][j] == 1 {
				// pat = 0x111, j=0の場合 dp[110]から組み合わせの数をもらってくる
				np := pat | (1 << j)
				dp[np] += dp[pat]
				dp[np] %= p
			}
		}
	}
	return dp[1<<n-1]
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextIntSlice(n)
	}
	//ans := SolveHonestly(n, a)
	ans := Solve(n, a)
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
