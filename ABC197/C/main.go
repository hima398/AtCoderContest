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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	//fmt.Println(ans)
	ans := 1 << 60
	max := 1 << (n - 1)
	for bits := 1; bits <= max; bits++ {
		xor := 0
		or := 0
		for i := 0; i <= n; i++ {
			if i < n {
				or |= a[i]
			}
			if i == n || bits&(1<<i) > 0 {
				xor ^= or
				or = 0
			}
		}
		ans = Min(ans, xor)
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
