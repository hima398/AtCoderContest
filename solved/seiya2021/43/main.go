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
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = nextString()
	}
	ans := 1 << (n + 1)
	nt := make([]int, n+1)
	nf := make([]int, n+1)
	nt[0] = ans / 2
	nf[0] = ans / 2
	for i := 1; i <= n; i++ {
		if s[i-1] == "AND" {
			nt[i] = nt[i-1] / 2
			nf[i] = ans - nt[i]
		} else {
			nf[i] = nf[i-1] / 2
			nt[i] = ans - nf[i]
		}
	}
	fmt.Println(nt[n])
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
