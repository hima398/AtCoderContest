package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func PrintSum(s [][]int) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

func SolveHonestly(n, k int, a, b []int) int {
	hmx, wmx := 0, 0
	for i := 0; i < n; i++ {
		hmx = Max(hmx, a[i])
		wmx = Max(wmx, b[i])
	}
	s := make([][]int, hmx+1)
	for i := 0; i <= hmx; i++ {
		s[i] = make([]int, wmx+1)
	}
	for idx := 0; idx < n; idx++ {
		for i := a[idx]; i <= a[idx]+k; i++ {
			for j := b[idx]; j <= b[idx]+k; j++ {
				if i >= 0 && i <= hmx && j >= 0 && j <= wmx {
					s[i][j]++
				}
			}
		}
	}
	ans := 0
	for i := 0; i <= hmx; i++ {
		for j := 0; j <= wmx; j++ {
			//fmt.Println(a[i], b[i], s[a[i]][b[i]])
			ans = Max(ans, s[i][j])
		}
	}
	//PrintSum(s)
	return ans
}

func Solve(n, k int, a, b []int) int {
	hmx, wmx := 0, 0
	for i := 0; i < n; i++ {
		hmx = Max(hmx, a[i])
		wmx = Max(wmx, b[i])
	}
	//mxa++
	//mxb++
	s := make([][]int, hmx+1)
	for i := 0; i <= hmx; i++ {
		s[i] = make([]int, wmx+1)
	}
	for i := 0; i < n; i++ {
		ri, rj := a[i]+k+1, b[i]+k+1
		s[a[i]][b[i]]++
		if ri <= hmx {
			s[ri][b[i]]--
			if rj <= wmx {
				s[ri][rj]++
			}
		}
		if rj <= wmx {
			s[a[i]][rj]--
		}
	}
	//横方向に累積和を計算していく
	for i := 0; i <= hmx; i++ {
		for j := 0; j < wmx; j++ {
			s[i][j+1] += s[i][j]
		}
	}
	//縦方向に累積和を計算していく
	for j := 0; j <= wmx; j++ {
		for i := 0; i < hmx; i++ {
			s[i+1][j] += s[i][j]
		}
	}

	ans := 0
	for i := 0; i <= hmx; i++ {
		for j := 0; j <= wmx; j++ {
			//fmt.Println(a[i], b[i], s[a[i]][b[i]])
			ans = Max(ans, s[i][j])
		}
	}
	//PrintSum(s)
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	a, b := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i] = nextInt(), nextInt()
	}
	//ans := SolveHonestly(n, k, a, b)
	ans := Solve(n, k, a, b)
	fmt.Println(ans)
	/*
	 */
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
