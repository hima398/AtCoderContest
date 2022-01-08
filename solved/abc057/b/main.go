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

	n, m := nextInt(), nextInt()
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i] = nextInt(), nextInt()
	}
	c := make([]int, m)
	d := make([]int, m)
	for i := 0; i < m; i++ {
		c[i], d[i] = nextInt(), nextInt()
	}
	var ans []int
	for i := 0; i < n; i++ {
		idx := n
		v := 4*int(1e8) + 1
		for j := 0; j < m; j++ {
			d := Abs(a[i]-c[j]) + Abs(b[i]-d[j])
			if v > d {
				idx = j
				v = d
			}
		}
		ans = append(ans, idx+1)
		//fmt.Println(ans)
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for _, v := range ans {
		fmt.Println(v)
	}
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
