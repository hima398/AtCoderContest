package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextString()
	t := nextString()
	ans := len(s)
	for i := 0; i < len(s)-len(t)+1; i++ {
		tmp := 0
		for j := 0; j < len(t); j++ {
			if s[i+j] != t[j] {
				tmp++
			}
		}
		ans = Min(ans, tmp)
	}
	fmt.Println(ans)
}
