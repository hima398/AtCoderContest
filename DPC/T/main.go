package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

//O(N*N!)になりn=10くらい程度までしか解けない見込み
func SolveHonestly(n int, s string) (ans int) {
	const p = int(1e9) + 7
	NextPermutation := func(x sort.Interface) bool {
		n := x.Len() - 1
		if n < 1 {
			return false
		}
		j := n - 1
		for ; !x.Less(j, j+1); j-- {
			if j == 0 {
				return false
			}
		}
		l := n
		for !x.Less(j, l) {
			l--
		}
		x.Swap(j, l)
		for k, l := j+1, n; k < l; {
			x.Swap(k, l)
			k++
			l--
		}
		return true
	}
	var a []int
	for i := 1; i <= n; i++ {
		a = append(a, i)
	}

	for {
		ok := true
		for i := 0; i < n-1; i++ {
			ok = ok && ((s[i] == '<' && a[i] < a[i+1]) || (s[i] == '>' && a[i] > a[i+1]))
		}
		if !NextPermutation(sort.IntSlice(a)) {
			break
		}
		if ok {
			ans++
			ans %= p
		}
	}
	return ans
}

func Solve(n int, s string) int {
	const p = int(1e9) + 7

	// sのi番目の不等号まで評価して、i番目の不等号の右辺の数より大きいものがjになる組み合わせ
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n+1)
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < n; i++ {
		// dp[i-1][j]までの和s[j]
		dps := make([]int, n)
		dps[0] = dp[i-1][0]
		for j := 1; j < n; j++ {
			dps[j] = dps[j-1] + dp[i-1][j]
			dps[j] %= p
		}
		//fmt.Println("s = ", dps)
		if s[i-1] == '<' {
			for j := 0; j < n-i; j++ {
				dp[i][j] += (dps[n-i] - dps[j] + p) % p
				dp[i][j] %= p
				// ここまでループを回すとO(N**3)になるので累積和を使ってO(N**2)にする
				//for k := j + 1; k < n-i+1; k++ {
				//	dp[i][j] += dp[i-1][k]
				//	dp[i][j] %= p
				//}
			}
		} else {
			for j := 0; j < n-i; j++ {
				dp[i][j] += dps[j]
				dp[i][j] %= p
				//for k := 0; k <= j; k++ {
				//	dp[i][j] += dp[i-1][k]
				//	dp[i][j] %= p
				//}
			}
		}

	}
	//for i := 0; i < n; i++ {
	//	fmt.Println(dp[i])
	//}

	return dp[n-1][0]
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	s := nextString()

	//ans := SolveHonestly(n, s)
	ans := Solve(n, s)
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
