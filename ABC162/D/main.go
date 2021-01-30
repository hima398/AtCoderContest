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
	m := make(map[string]int)
	for i := 0; i < n; i++ {
		m[string(s[i])]++
	}
	ans := m["R"] * m["G"] * m["B"]
	for i := 0; i < n-1; i++ {
		for j := i; j < n; j++ {
			k := 2*j - i
			if k >= n {
				continue
			}
			if s[k] != s[j] && s[j] != s[i] && s[k] != s[i] {
				ans--
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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
