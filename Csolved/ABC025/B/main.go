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

	dir := map[string]int{"East": 1, "West": -1}

	n, a, b := nextInt(), nextInt(), nextInt()
	p := 0
	for i := 0; i < n; i++ {
		s, d := nextString(), nextInt()
		var w int
		if d < a {
			w = a
		} else if d > b {
			w = b
		} else {
			w = d
		}
		p += dir[s] * w
	}
	if p > 0 {
		fmt.Println("East", p)
	} else if p < 0 {
		fmt.Println("West", Abs(p))
	} else {
		fmt.Println(0)
	}
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
