package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

// N=30
//
func Solve1(n, w int, sv, sw []int) int {
	type Load struct {
		v, w int
	}
	hm, lm := 0xffff0000, 0xffff
	pat := (1 << n) - 1
	hpat, lpat := pat&hm, pat&lm
	hpat = hpat >> 16
	//fmt.Printf("%16b\n", lpat)
	//fmt.Printf("%16b\n", hpat)
	var h, l []Load
	for i := 0; i <= lpat; i++ {
		tv, tw := 0, 0
		for j := 0; j < Min(n, 16); j++ {
			if (i>>j)&1 == 1 {
				tv += sv[j]
				tw += sw[j]
			}
		}
		l = append(l, Load{tv, tw})
	}
	for i := 0; i <= hpat; i++ {
		tv, tw := 0, 0
		for j := 0; j < Min(n-16, 16); j++ {
			if (i>>j)&1 == 1 {
				tv += sv[j+16]
				tw += sw[j+16]
			}
		}
		h = append(h, Load{tv, tw})
	}
	sort.Slice(l, func(i, j int) bool {
		return l[i].w < l[j].w
	})
	sort.Slice(h, func(i, j int) bool {
		return h[i].w < h[j].w
	})
	//fmt.Println(l)
	//fmt.Println(h)
	ans := 0
	for i := 0; i < len(l); i++ {
		if l[i].w > w {
			break
		}
		idx := sort.Search(len(h), func(j int) bool {
			return l[i].w+h[j].w > w
		})
		//if idx >= len(h) {
		//	fmt.Println(idx, w, l[i].w+h[idx-1].w)
		//} else {
		//	fmt.Println(idx, w, l[i].w+h[idx-1].w, l[i].w+h[idx].w)
		//}
		for j := 0; j < idx; j++ {
			ans = Max(ans, l[i].v+h[j].v)
		}
	}
	return ans
}

// wi <= 1000
func Solve2(n, w int, sv, sw []int) int {
	// dp[0][0] = 0, dp[1][9] = 15, dp[2][6] = 10 dp[2][9]=15, dp[3][6] = 10
	const INF = 1 << 60
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, w+1)
		for j := 0; j <= w; j++ {
			dp[i][j] = -INF
		}
	}
	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := 0; j <= w; j++ {
			dp[i][j] = dp[i-1][j]
			if j-sw[i-1] < 0 {
				continue
			}
			dp[i][j] = Max(dp[i][j], dp[i-1][j-sw[i-1]]+sv[i-1])
		}
	}
	ans := 0
	for i := 0; i <= w; i++ {
		ans = Max(ans, dp[n][i])
	}
	return ans
}

// vi <= 1000
func Solve3(n, w int, sv, sw []int) int {
	const INF = 1 << 60
	vv := 0
	for i := 0; i < n; i++ {
		vv += sv[i]
	}
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, vv+1)
		for j := 0; j <= vv; j++ {
			dp[i][j] = INF
		}
	}
	dp[0][0] = 0
	for i := 1; i <= n; i++ {
		for j := 0; j <= vv; j++ {
			dp[i][j] = dp[i-1][j]
			if j-sv[i-1] < 0 {
				continue
			}
			dp[i][j] = Min(dp[i][j], dp[i-1][j-sv[i-1]]+sw[i-1])
		}
	}
	ans := 0
	for i := 0; i <= vv; i++ {
		if dp[n][i] > w {
			continue
		}
		ans = i
	}
	return ans
}

func Solve(n, w int, sv, sw []int) int {
	vmax, wmax := 0, 0
	for i := 0; i < n; i++ {
		vmax += sv[i]
		wmax += sw[i]
	}
	if wmax <= 200*1000 {
		return Solve2(n, w, sv, sw)

	} else if vmax <= 200*1000 {
		return Solve3(n, w, sv, sw)

	} else {
		return Solve1(n, w, sv, sw)
	}
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, w := nextInt(), nextInt()
	sv, sw := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		sv[i], sw[i] = nextInt(), nextInt()
	}
	ans := Solve(n, w, sv, sw)
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
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
