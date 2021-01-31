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

	n, a, b := nextInt(), nextInt(), nextInt()
	ans := n / (a + b)
	ans *= a
	ans += Min(n%(a+b), a)
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
