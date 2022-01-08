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
	s := nextString()
	var ans int
	for i := 1; i < n-1; i++ {
		m1 := make(map[byte]bool)
		for j := 0; j < i; j++ {
			m1[s[j]] = true
		}
		m2 := make(map[byte]bool)
		for j := i; j < n; j++ {
			m2[s[j]] = true
		}
		var cnt int
		for k := 0; k < 26; k++ {
			if m1[byte('a'+k)] && m2[byte('a'+k)] {
				cnt++
			}
		}
		ans = Max(ans, cnt)
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
