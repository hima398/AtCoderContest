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
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = nextInt()
	}
	ans := 0
	for i := 1; i < n-1; i++ {
		if (p[i] > p[i-1] && p[i] < p[i+1]) || (p[i] > p[i+1] && p[i] < p[i-1]) {
			ans++
		}
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
