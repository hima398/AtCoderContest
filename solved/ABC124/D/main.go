package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Solve(k int, dig []int) int {
	add := 2*k + 1
	l := 0
	r := 0
	sum := 0 // [l, r)のsum
	ans := 0
	n := len(dig)
	for i := 0; i < n; i += 2 {
		nextL := i
		nextR := Min(i+add, n)
		for nextL > l {
			sum -= dig[l]
			l++
		}
		for nextR > r {
			sum += dig[r]
			r++
		}
		//fmt.Println(i, sum)
		ans = Max(ans, sum)
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(), nextInt()
	s := nextString()
	var dig []int
	now := byte(1)
	cnt := 0
	for i := 0; i < n; i++ {
		if s[i] == '0'+now {
			cnt++
		} else {
			dig = append(dig, cnt)
			now ^= 1
			cnt = 1
		}
	}
	// cntが0ではないチェックがいるかも？
	dig = append(dig, cnt)
	// 長さが奇数になるように調整
	if len(dig)%2 == 0 {
		dig = append(dig, 0)
	}
	// 11101010110011
	// 3 1 1 1 1 1 2 2
	fmt.Println(Solve(k, dig))
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
