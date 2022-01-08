package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	a := "ACGT"
	s := nextString()
	sub, ans := 0, 0
	for _, v := range s {
		if strings.Contains(a, string(v)) {
			sub++
		} else {
			ans = Max(ans, sub)
			sub = 0
		}
	}
	ans = Max(ans, sub)
	fmt.Println(ans)
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
