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

	n, m := nextInt(), nextInt()
	x := make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = -1
	}
	for i := 0; i < m; i++ {
		j, c := nextInt(), nextInt()
		j--
		// 矛盾
		if x[j] != -1 && x[j] != c {
			fmt.Println(-1)
			return
		}
		x[j] = c
	}
	if n == 1 {
		if x[0] == -1 {
			fmt.Println(0)
		} else {
			fmt.Println(x[0])
		}
		return
	}
	// n == 2 or 3
	if x[0] == 0 {
		fmt.Println(-1)
		return
	} else if x[0] == -1 {
		x[0] = 1
	}
	for i := 1; i < n; i++ {
		if x[i] == -1 {
			x[i] = 0
		}
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%d", x[i])
	}
	fmt.Println("")
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
