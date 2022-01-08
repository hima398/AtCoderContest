package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveHonestly(n int, a []int) int {
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		s[i] += s[i-1] + a[i-1]
	}
	fmt.Println(s)

	ans := 1 << 60
	t := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sp := 0
		op := 0
		// s[i]を正にする場合
		for j := 1; j <= n; j++ {
			if j%2 == i%2 {
				if s[j]+op <= 0 {
					sp += Abs(s[j]) + 1
					op += Abs(s[j]) + 1
				}
			} else {
				if s[j]+op >= 0 {
					sp += Abs(s[j]) + 1
					op -= s[j] + 1
				}
			}
		}
		sm := 0
		om := 0
		// s[i]を負にする場合
		for j := 1; j <= n; j++ {
			if j%2 == i%2 {
				if s[j]+om >= 0 {
					sm += Abs(s[j]) + 1
					om -= s[j] + 1
				}
			} else {
				if s[j]+om <= 0 {
					sm += Abs(s[j]) + 1
					om += Abs(s[j]) + 1
				}
			}
		}
		t[i] = Min(sp, sm)
		ans = Min(ans, t[i])
	}
	//ここでわかるのが、s[i]のどこを固定しても
	//手数の最小は同じ値になりそう
	//なのでs[i]を正にするか負にするかだけ決めて
	//a[i]に行った操作をoffsetに残せば良さそう
	fmt.Println(t)
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := nextIntSlice(n)
	//fmt.Println(SolveHonestly(n, a))
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		s[i] += s[i-1] + a[i-1]
	}

	//s[1]を正にする場合の手数
	sp := 0
	//s[i]を更新すると、i+1以降に加える必要がある値
	op := 0
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			if s[i]+op <= 0 {
				sp += Abs(s[i]+op) + 1
				op += Abs(s[i]+op) + 1
			}
		} else {
			if s[i]+op >= 0 {
				sp += Abs(s[i]+op) + 1
				op -= Abs(s[i]+op) + 1
			}
		}
	}
	//s[1]を負にする場合の手数
	sm := 0
	//s[i]を更新すると、i+1以降に加える必要がある値
	om := 0
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			if s[i]+om >= 0 {
				sm += Abs(s[i]+om) + 1
				om -= Abs(s[i]+om) + 1
			}
		} else {
			if s[i]+om <= 0 {
				sm += Abs(s[i]+om) + 1
				om += Abs(s[i]+om) + 1
			}
		}
	}
	ans := Min(sp, sm)
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
