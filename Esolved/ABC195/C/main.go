package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveCommentary(n int) int {
	ans := 0
	if n >= int(1e3) {
		ans += n - (int(1e3) - 1)
	}
	if n >= int(1e6) {
		ans += n - (int(1e6) - 1)
	}
	if n >= int(1e9) {
		ans += n - (int(1e9) - 1)
	}
	if n >= int(1e12) {
		ans += n - (int(1e12) - 1)
	}
	if n >= int(1e15) {
		ans += n - (int(1e15) - 1)
	}
	return ans
}

func FirstSolve(n int) int {
	ans := 0
	if n < int(1e3) {
		return ans
	}
	ans += Min(n, int(1e6)-1) - int(1e3) + 1
	if n < int(1e6) {
		return ans
	}
	ans += (Min(n, int(1e9)-1) - int(1e6) + 1) * 2
	if n < int(1e9) {
		return ans
	}
	ans += (Min(n, int(1e12)-1) - int(1e9) + 1) * 3
	if n < int(1e12) {
		return ans
	}
	ans += (Min(n, int(1e15)-1) - int(1e12) + 1) * 4
	if n < int(1e15) {
		return ans
	}
	ans += 5
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	//fmt.Println(SolveCommentary(n))
	fmt.Println(FirstSolve(n))
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
