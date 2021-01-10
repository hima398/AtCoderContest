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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a, b, c, d := nextInt(), nextInt(), nextInt(), nextInt()
	ans := -9223372036854775808
	ans = Max(ans, b*d)
	ans = Max(ans, b*c)
	ans = Max(ans, a*c)
	ans = Max(ans, a*d)
	fmt.Println(ans)
}
