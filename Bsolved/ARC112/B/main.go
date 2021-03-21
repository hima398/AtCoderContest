package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func FirstCode(b, c int) int {
	// 色々と試行錯誤したけどもこのやり方では正解にたどり着けていない
	// -1できる回数
	var minus int

	minus = c / 2
	var ans int
	if b == 0 {
		if c%2 == 1 {
			// T1
			ans = 2*minus + 1
		} else {
			// E2
			ans = 2 * minus
		}
		fmt.Println(ans)
		return ans
	}
	if b < 0 || (b > 0 && b-minus > 0) {
		if c%2 == 1 {
			// T2
			ans = 2*minus + 1
		} else {
			// E4
			ans = 2 * minus
		}
		ans += Min(2*minus, 2*Abs(b))
	} else {
		// b > 0 && b-minus <= 0 {
		if c%2 == 1 {
			ans = 2*minus + 1
		} else {
			ans = 2 * minus
		}
	}
	return ans
}

func ComputeRange(lb, lc int) (int, int) {
	n := lc / 2
	if lc%2 == 1 {
		return -lb - n, -lb + n
	} else {
		if lc == 0 {
			return lb, lb
		} else {
			return lb - n, lb + n - 1
		}
	}
}
func SolveCommentary(lb, lc int) int {
	a, b := ComputeRange(lb, lc)
	c, d := ComputeRange(lb, lc-1)
	return (b - a + 1) + (d - c + 1) - Max(0, Min(b, d)-Max(a, c)+1)
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	lb, lc := nextInt(), nextInt()
	fmt.Println(SolveCommentary(lb, lc))
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
