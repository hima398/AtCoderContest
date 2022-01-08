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
	ac := make([]bool, n+1)
	wa := make([]int, n+1)
	nac := 0
	np := 0
	for i := 0; i < m; i++ {
		p, s := nextInt(), nextString()
		if ac[p] {
			continue
		}
		// not ac
		if s == "AC" {
			nac++
			np += wa[p]
			ac[p] = true
		}
		if s == "WA" {
			wa[p]++
		}
	}
	fmt.Println(nac, np)
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
