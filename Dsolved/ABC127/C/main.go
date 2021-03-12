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
	c := make([]int, n+1)
	for i := 0; i < m; i++ {
		l, r := nextInt(), nextInt()
		l--
		c[l]++
		c[r]--
	}
	//fmt.Println(c)
	ans := 0
	for i := 0; i < n; i++ {
		if c[i] == m {
			ans++
		}
		c[i+1] += c[i]
	}
	//fmt.Println(c)
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
