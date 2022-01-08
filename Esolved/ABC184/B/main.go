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

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func minus(x, y int) int {
	r := x - y
	if r < 0 {
		return 0
	}
	return r
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, x := nextInt(), nextInt()
	s := nextString()
	result := x
	for i := 0; i < n; i++ {
		if s[i] == 'o' {
			result++
		} else {
			// 'x'
			result = minus(result, 1)
		}
	}
	fmt.Println(result)
}
