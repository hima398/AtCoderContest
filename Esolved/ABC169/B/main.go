package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := make([]int, n)
	const max = 1000000000000000000
	for i := 0; i < n; i++ {
		a[i] = nextInt()
		if a[i] == 0 {
			fmt.Println(0)
			return
		}
	}
	ans := 1
	for i := 0; i < n; i++ {
		if max/ans >= a[i] {
			ans *= a[i]
		} else {
			fmt.Println(-1)
			return
		}
	}
	fmt.Println(ans)
}
