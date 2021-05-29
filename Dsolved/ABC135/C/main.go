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
	for i := 0; i < n+1; i++ {
		a[i] = nextInt()
	}
	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = nextInt()
	}
	ans := 0
	for i := 0; i < n; i++ {
		if a[i] <= b[i] {
			ans += a[i]
			b[i] -= a[i]
			a[i] = 0
		} else {
			ans += b[i]
			continue
		}
		if a[i+1] <= b[i] {
			ans += a[i+1]
			a[i+1] = 0
		} else {
			ans += b[i]
			a[i+1] -= b[i]
		}
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
