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
	a, b, c := nextInt(), nextInt(), nextInt()
	const max = int(1e4)
	ans := max
	for i := 0; i < max; i++ {
		for j := 0; j < max; j++ {
			r := n - (a*i + b*j)
			if r >= 0 && r%c == 0 {
				ans = Min(ans, i+j+r/c)
			}
		}
	}
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
