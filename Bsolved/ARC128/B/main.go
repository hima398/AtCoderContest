package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Solve(r, g, b int) int {
	const INF = int(1e8) + 1
	ar := INF
	if g%3 == b%3 {
		ar = Max(g, b)
	}
	ag := INF
	if r%3 == b%3 {
		ag = Max(r, b)
	}
	ab := INF
	if r%3 == g%3 {
		ab = Max(r, g)
	}
	ans := Min(ar, Min(ag, ab))
	if ans == INF {
		ans = -1
	}
	return ans
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	t := nextInt()
	var ans []int
	for i := 0; i < t; i++ {
		r, g, b := nextInt(), nextInt(), nextInt()
		ans = append(ans, Solve(r, g, b))
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, v := range ans {
		fmt.Fprintln(out, v)
	}
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
