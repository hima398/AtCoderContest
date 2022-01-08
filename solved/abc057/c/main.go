package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func ComputeDivisor(n int) []int {
	var ret []int
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			ret = append(ret, i)
		}
	}
	return ret
}

func F(a, b int) int {
	sa, sb := strconv.Itoa(a), strconv.Itoa(b)
	return Max(len(sa), len(sb))
}
func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	d := ComputeDivisor(n)
	//fmt.Println(d)
	ans := 11
	for _, v := range d {
		ans = Min(ans, F(v, n/v))
	}
	fmt.Println(ans)
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
