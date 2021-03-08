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
	s := make([]int, n+1)
	s2 := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = nextInt()
		s[i] = s[i-1] + a[i]
		s2[i] = s2[i-1] + (a[i] * a[i])
	}
	sum := 0
	for i := 2; i <= n; i++ {
		sum += (i-1)*a[i]*a[i] - 2*a[i]*s[i-1] + s2[i-1]
	}
	fmt.Println(sum)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
