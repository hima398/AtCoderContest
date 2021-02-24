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

	n, k := nextInt(), nextInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	s := 1
	for i := 0; i < k; i++ {
		s *= a[i]
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := k; i < n; i++ {
		if a[i-k] < a[i] {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
