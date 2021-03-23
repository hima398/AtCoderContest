package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	s := nextString()
	a1 := 0
	// 0101....と塗り分ける
	for i := 0; i < len(s); i++ {
		if i%2 == 0 && s[i] != '0' {
			a1++
		} else if i%2 == 1 && s[i] != '1' {
			a1++
		}
	}
	a2 := 0
	// 1010....と塗り分ける
	for i := 0; i < len(s); i++ {
		if i%2 == 0 && s[i] != '1' {
			a2++
		} else if i%2 == 1 && s[i] != '0' {
			a2++
		}
	}
	ans := Min(a1, a2)
	fmt.Println(ans)
}

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
