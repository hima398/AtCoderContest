package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func FirstSolve(n, a, b int, v []int) (float64, int) {

	dp := make([][]float64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]float64, n+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i][j] = math.Max(dp[i-1][j], (dp[i-1][j-1]*float64(j-1)+float64(v[i-1]))/float64(j))
		}
	}
	ans1 := 0.0
	for i := a; i <= b; i++ {
		ans1 = math.Max(ans1, dp[n][i])
	}
	return ans1, 0
}

func Solve(n, a, b int, v []int) (float64, int) {
	m1 := make(map[int]int)
	for _, vi := range v {
		m1[vi]++
	}
	sort.Slice(v, func(i, j int) bool {
		return v[i] > v[j]
	})
	s := 0
	var v2 []int
	m2 := make(map[int]int)
	for i := 0; i < a; i++ {
		s += v[i]
		m2[v[i]]++
		v2 = append(v2, v[i])
	}
	sort.Slice(v2, func(i, j int) bool {
		return v2[i] > v2[j]
	})
	ans1 := float64(s) / float64(a)
	ans2 := 0
	//fmt.Println(m1)
	//fmt.Println(m2)
	if len(m2) == 1 {
		// ソート済みのv0をa個取った平均が最大の場合はa〜b個くらいのの選び方で決まる
		for i := a; i <= Min(b, m1[v2[0]]); i++ {
			//for i := a; i <= b; i++ {
			s := 1
			s *= Combination(m1[v2[0]], i)
			ans2 += s
		}
	} else {
		// ソート済みのvから複数種類の数字の平均の場合は一番小さい数の選び方で決まる
		k := v2[len(v2)-1]
		ans2 = Combination(m1[k], m2[k])
	}

	return ans1, ans2
}

func Estimate() int {
	res := 0
	for i := 1; i <= 50; i++ {
		res += Combination(50, i)
	}
	return res
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, a, b := nextInt(), nextInt(), nextInt()
	v := nextIntSlice(n)

	//ans1, ans2 := FirstSolve(n, a, b, v)
	ans1, ans2 := Solve(n, a, b, v)
	fmt.Println(ans1)
	fmt.Println(ans2)
	//fmt.Println(Estimate())
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

func Permutation(N, K int) int {
	v := 1
	if 0 < K && K <= N {
		for i := 0; i < K; i++ {
			v *= (N - i)
		}
	} else if K > N {
		v = 0
	}
	return v
}

func Factional(N int) int {
	return Permutation(N, N-1)
}

func Combination(N, K int) int {
	if K == 0 {
		return 1
	}
	if K == 1 {
		return N
	}
	return Combination(N, K-1) * (N + 1 - K) / K
}

type Comb struct {
	n, p int
	fac  []int // Factional(i) mod p
	finv []int // 1/Factional(i) mod p
	inv  []int // 1/i mod p
}

func NewCombination(n, p int) *Comb {
	c := new(Comb)
	c.n = n
	c.p = p
	c.fac = make([]int, n+1)
	c.finv = make([]int, n+1)
	c.inv = make([]int, n+1)

	c.fac[0] = 1
	c.fac[1] = 1
	c.finv[0] = 1
	c.finv[1] = 1
	c.inv[1] = 1
	for i := 2; i <= n; i++ {
		c.fac[i] = c.fac[i-1] * i % p
		c.inv[i] = p - c.inv[p%i]*(p/i)%p
		c.finv[i] = c.finv[i-1] * c.inv[i] % p
	}
	return c
}

func (c *Comb) Factional(x int) int {
	return c.fac[x]
}

func (c *Comb) Combination(n, k int) int {
	if n < k {
		return 0
	}
	if n < 0 || k < 0 {
		return 0
	}
	ret := c.fac[n] * c.finv[k]
	ret %= c.p
	ret *= c.finv[n-k]
	ret %= c.p
	return ret
}
