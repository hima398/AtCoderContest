package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

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

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	x, n := nextInt(), nextInt()
	m := make(map[int]bool)
	for i := 0; i < n; i++ {
		p := nextInt()
		m[p] = true
	}
	min := 100
	ans := 0
	for i := 0; i <= 101; i++ {
		if m[i] {
			continue
		}
		d := Abs(x - i)
		if min > d {
			min = d
			ans = i
		}
	}
	fmt.Println(ans)
}
