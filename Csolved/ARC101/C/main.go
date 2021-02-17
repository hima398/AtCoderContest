package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveHonestly(n, k int) int {
	x := make([]int, n)
	x[0] = nextInt()
	for i := 1; i < n; i++ {
		x[i] = nextInt()
	}
	ans := int(1e14) // > |xi|(1e8) * nmax(1e5)
	for i := 0; i <= n-k; i++ {
		px := x[i : i+k]
		ps := 0
		pre := px[0]
		for _, v := range px {
			ps += Abs(v - pre)
			pre = v
		}
		d := Min(Abs(x[i+k-1])+ps, Abs(x[i])+ps)
		ans = Min(ans, d)
	}
	return ans
}

func Solve(n, k int) int {
	x := make([]int, n)
	s := make([]int, n)
	x[0] = nextInt()
	s[0] = 0
	for i := 1; i < n; i++ {
		x[i] = nextInt()
		s[i] = s[i-1] + Abs(x[i]-x[i-1])
	}
	//fmt.Println(x)
	//fmt.Println(s)
	ans := int(1e14) // > |xi|(1e8) * nmax(1e5)
	for i := 0; i <= n-k; i++ {
		//px := x[i : i+k]
		//fmt.Println(px)
		//fmt.Println(i)
		ps := s[i+k-1] - s[i]
		//fmt.Println(ps, s[i+k-1], s[i])
		d := Min(Abs(x[i+k-1])+ps, Abs(x[i])+ps)
		ans = Min(ans, d)
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	//fmt.Println(SolveHonestly(n, k))
	fmt.Println(Solve(n, k))
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
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
