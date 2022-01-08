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
	n, l := nextInt(), nextInt()
	a := make([]int, n)
	hasZero := false
	s := 0
	for i := 0; i < n; i++ {
		a[i] = l + i
		if a[i] == 0 {
			hasZero = true
		}
		s += a[i]
	}
	if hasZero {
		fmt.Println(s)
	} else {
		if l > 0 {
			fmt.Println(s - a[0])
		} else {
			fmt.Println(s - a[n-1])
		}
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
