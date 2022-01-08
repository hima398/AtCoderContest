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
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = nextString()
	}
	var t, a int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if s[i][j] == 'R' {
				t++
			} else if s[i][j] == 'B' {
				a++
			}
		}
	}
	if t > a {
		fmt.Println("TAKAHASHI")
	} else if t < a {
		fmt.Println("AOKI")
	} else {
		fmt.Println("DRAW")
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
