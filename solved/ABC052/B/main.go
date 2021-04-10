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

	_ = nextInt()
	s := nextString()
	x := 0
	ans := x
	for i := range s {
		if s[i] == 'I' {
			x++
		} else {
			x--
		}
		ans = Max(ans, x)
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
