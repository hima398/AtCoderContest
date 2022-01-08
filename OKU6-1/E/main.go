package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveHonestly(n int, a []int) (ans int) {
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			score := a[i] + a[j] - ((a[i] & a[j]) ^ (a[i] | a[j]))
			ans = Max(ans, score)
		}
	}
	return ans
}

func Solve(n int, a []int) (ans int) {
	for k := 31; k >= 0; k-- {
		cnt := 0
		for i := 0; i < n; i++ {
			bit := 1 << k
			if a[i]&ans == ans && a[i]&bit > 0 {
				cnt++
			}
			if cnt >= 2 {
				ans |= bit
			}
		}
	}
	return 2 * ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := nextIntSlice(n)

	//fmt.Println(SolveHonestly(n, a))
	fmt.Println(Solve(n, a))
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
