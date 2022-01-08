package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

// 近傍のみ検索でWA x 14、近傍とコスト最小の組み合わせを可能な限り全探索で
func FirstSolve(h, w, c int, a [][]int) int {
	const INF = 1 << 60

	//線路を最小にする(上、左、下、右)
	di := []int{-1, 0, 1, 0}
	dj := []int{0, -1, 0, 1}
	ans2 := INF
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			for k := 0; k < 4; k++ {
				ni := i + di[k]
				nj := j + dj[k]
				if ni >= 0 && ni < h && nj >= 0 && nj < w {
					cost := a[i][j] + a[ni][nj] + c*(Abs(i-ni)+Abs(j-nj))
					ans2 = Min(ans2, cost)
				}
			}
		}
	}

	//駅のコストを最小にする組み合わせを可能な限り検索
	type Pos struct {
		i, j, c int
	}
	ans1 := INF
	var ps []Pos
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ps = append(ps, Pos{i, j, a[i][j]})
		}
	}
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].c < ps[j].c
	})
	for i := 0; i < Min(len(ps), 30000)-1; i++ {
		for j := i + 1; j < Min(len(ps), 30000); j++ {
			cost := ps[i].c + ps[j].c + c*(Abs(ps[i].i-ps[j].i)+Abs(ps[i].j-ps[j].j))
			ans2 = Min(ans2, cost)
		}
	}
	ans := Min(ans1, ans2)
	return ans
}

func SecondSolve(h, w, c int, a [][]int) int {
	const INF = 1 << 60
	type Pos struct {
		i, j, min, cost int // 最も右下にいる最小の駅のi, j, 駅のコスト、最小コスト
	}
	dp := make([][]Pos, h)
	for i := 0; i < h; i++ {
		dp[i] = make([]Pos, w)
	}
	dp[0][0] = Pos{0, 0, a[0][0], INF}
	for j := 1; j < w; j++ {
		if dp[0][j-1].min >= a[0][j] {
			dp[0][j].j = j
			dp[0][j].min = a[0][j]
		} else {
			dp[0][j].i = dp[0][j-1].i
			dp[0][j].j = dp[0][j-1].j
			dp[0][j].min = dp[0][j-1].min
		}
		//隣同士
		cost1 := a[0][j-1] + a[0][j] + c
		//1つ手前の最小の駅との距離
		cost2 := dp[0][j-1].min + a[0][j] + c*(Abs(dp[0][j-1].j-j))
		dp[0][j].cost = Min(dp[0][j-1].cost, Min(cost1, cost2))
	}
	for i := 1; i < h; i++ {
		if dp[i-1][0].min >= a[i][0] {
			dp[i][0].i = i
			dp[i][0].min = a[i][0]
		} else {
			dp[i][0].i = dp[i-1][0].i
			dp[i][0].j = dp[i-1][0].j
			dp[i][0].min = dp[i-1][0].min
		}
		//隣同士
		cost1 := a[i-1][0] + a[i][0] + c
		//1つ手前の最小の駅との距離
		cost2 := dp[i-1][0].min + a[i][0] + c*(Abs(dp[i-1][0].i-i))
		dp[i][0].cost = Min(dp[i-1][0].cost, Min(cost1, cost2))
	}
	for i := 1; i < h; i++ {
		for j := 1; j < w; j++ {
			if dp[i][j-1].min >= a[i][j] {
				dp[i][j].i = i
				dp[i][j].j = j
				dp[i][j].min = a[i][j]
			} else {
				dp[i][j].i = dp[i][j-1].i
				dp[i][j].j = dp[i][j-1].j
				dp[i][j].min = dp[i][j-1].min
			}
			if dp[i-1][j].min >= a[i][j] {
				dp[i][j].i = i
				dp[i][j].j = j
				dp[i][j].min = a[i][j]
			} else {
				dp[i][j].i = dp[i][j-1].i
				dp[i][j].j = dp[i][j-1].j
				dp[i][j].min = dp[i][j-1].min
			}
			//隣同士
			cost1 := Min(a[i-1][j]+a[i][j]+c, a[i][j-1]+a[i][j]+c)
			//1つ手前の最小の駅との距離
			min21 := dp[i-1][j].min + a[i][j] + c*(Abs(dp[i-1][j].i-i)+Abs(dp[i-1][j].j-j))
			min22 := dp[i][j-1].min + a[i][j] + c*(Abs(dp[i][j-1].i-i)+Abs(dp[i][j-1].j-j))
			cost2 := Min(min21, min22)
			dp[i][j].cost = Min(dp[i-1][j].cost, Min(dp[i][j-1].cost, Min(cost1, cost2)))
		}
	}
	/*
		PrintDp := func(dp [][]Pos) {
			for i := 0; i < len(dp); i++ {
				fmt.Println(dp[i])
			}
		}
	*/
	//PrintDp(dp)
	return dp[h-1][w-1].cost

}

func SolveCommentary(h, w, c int, a [][]int) int {
	const INF = int(1e15) + 1 // 1000 * 1000 * 10**9より大きい数

	var a2 [2][][]int
	for k := 0; k < 2; k++ {
		a2[k] = make([][]int, h)
		for i := 0; i < h; i++ {
			a2[k][i] = make([]int, w)
			for j := 0; j < w; j++ {
				if k == 0 {
					a2[k][i][j] = a[i][j]
				} else if k == 1 {
					a2[k][i][j] = a[i][w-j-1]
				}
			}
		}
	}
	var dps [2][][]int
	var dpt [2][][]int
	for k := 0; k < 2; k++ {
		dps[k] = make([][]int, h)
		dpt[k] = make([][]int, h)
		for i := 0; i < h; i++ {
			dps[k][i] = make([]int, w)
			dpt[k][i] = make([]int, w)
			for j := 0; j < w; j++ {
				dps[k][i][j] = INF
				dpt[k][i][j] = INF
			}
		}
	}
	dps[0][0][0] = a2[0][0][0]
	dps[1][0][0] = a2[1][0][0]
	for k := 0; k < 2; k++ {
		for i := 1; i < h; i++ {
			dps[k][i][0] = Min(a2[k][i][0], dps[k][i-1][0]+c)
		}
		for j := 1; j < w; j++ {
			dps[k][0][j] = Min(a2[k][0][j], dps[k][0][j-1]+c)
		}
		for i := 1; i < h; i++ {
			for j := 1; j < w; j++ {
				dps[k][i][j] = Min(a2[k][i][j], Min(dps[k][i-1][j], dps[k][i][j-1])+c)
			}
		}
	}

	for k := 0; k < 2; k++ {
		for i := 1; i < h; i++ {
			dpt[k][i][0] = dps[k][i-1][0] + c + a2[k][i][0]
		}
		for j := 1; j < w; j++ {
			dpt[k][0][j] = dps[k][0][j-1] + c + a2[k][0][j]
		}
		for i := 1; i < h; i++ {
			for j := 1; j < w; j++ {
				dpt[k][i][j] = Min(dps[k][i-1][j], dps[k][i][j-1]) + c + a2[k][i][j]
			}
		}
	}

	ans := INF
	for k := 0; k < 2; k++ {
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				ans = Min(ans, dpt[k][i][j])
			}
		}
	}

	return ans

}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	h, w, c := nextInt(), nextInt(), nextInt()
	a := make([][]int, h)
	for i := 0; i < h; i++ {
		a[i] = nextIntSlice(w)
	}

	ans := FirstSolve(h, w, c, a)
	//ans := SolveCommentary(h, w, c, a)
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
