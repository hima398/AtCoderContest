package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a, b := nextInt(), nextInt()
	var ans []int
	max, min := Max(a, b), Min(a, b)
	for i := 1; i <= max; i++ {
		ans = append(ans, i)
	}
	smax := max * (max + 1) / 2
	smin := 0
	for i := 1; i < min; i++ {
		smin += i
		ans = append(ans, -i)
	}
	ans = append(ans, -(smax - smin))
	if a < b {
		for i := range ans {
			ans[i] = -ans[i]
		}
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, v := range ans {
		fmt.Fprintf(out, "%d ", v)
	}
	fmt.Fprintln(out, "")
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
