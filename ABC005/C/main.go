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

	t := nextInt()
	n := nextInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	m := nextInt()
	b := make([]int, m)
	for i := 0; i < m; i++ {
		b[i] = nextInt()
	}
	ti := 0
	ci := 0
	for ci < m {
		if a[ti] > b[ci] {
			break
		}
		for ti < n {
			if a[ti]+t >= b[ci] {
				//fmt.Println(ti, ci)
				ti++
				ci++
				break
			}
			ti++
		}
		if ti == n {
			break
		}
	}
	if ci == m {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
