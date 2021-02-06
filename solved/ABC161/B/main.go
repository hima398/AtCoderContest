package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const Mod = 1000000007

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	a := make([]int, n)
	s := 0
	for i := 0; i < n; i++ {
		a[i] = nextInt()
		s += a[i]
	}
	c := 0
	for i := 0; i < n; i++ {
		if 4*m*a[i] >= s {
			c++
		}
	}
	if c >= m {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
