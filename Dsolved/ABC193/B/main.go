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
	ans := 1000000001
	for i := 1; i <= n; i++ {
		a, p, x := nextInt(), nextInt(), nextInt()
		if x-a <= 0 {
			continue
		}
		ans = Min(ans, p)
	}
	if ans == 1000000001 {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
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
