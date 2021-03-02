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
	b := make([]int, n+1)
	c := make([]int, n)
	for i := 1; i <= n; i++ {
		a[i] = nextInt()
	}
	for i := 1; i <= n; i++ {
		b[i] = nextInt()
	}
	for i := 1; i <= n-1; i++ {
		c[i] = nextInt()
	}
	ans := 0
	pre := -1
	for i := 1; i <= n; i++ {
		ans += b[a[i]]
		if a[i]-1 == pre {
			ans += c[pre]
		}
		pre = a[i]
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
