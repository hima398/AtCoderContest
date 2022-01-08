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

	n, x := nextInt(), nextInt()
	l := nextIntSlice(n)
	d := 0
	ans := 1
	for i := 1; i <= n; i++ {
		if d+l[i-1] <= x {
			d += l[i-1]
			ans++
		} else {
			fmt.Println(ans)
			return
		}
	}
	fmt.Println(n + 1)
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
