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

	n := nextInt()
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = nextInt()
	}
	mi := make(map[int]int)
	ans := 0
	for j := 1; j <= n; j++ {
		r := j - a[j]
		l := j + a[j]
		ans += mi[r]
		mi[l]++
	}
	//fmt.Println(mi)
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
