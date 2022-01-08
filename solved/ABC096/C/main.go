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

	h, w := nextInt(), nextInt()
	s := make([]string, h)
	for i := 0; i < h; i++ {
		s[i] = nextString()
	}
	di := []int{-1, 0, 1, 0}
	dj := []int{0, -1, 0, 1}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == '#' {
				ok := false
				for k := 0; k < 4; k++ {
					ni, nj := i+di[k], j+dj[k]
					if ni >= 0 && ni < h && nj >= 0 && nj < w {
						ok = ok || s[ni][nj] == '#'
					}
				}
				if !ok {
					fmt.Println("No")
					return
				}
			}
		}
	}
	fmt.Println("Yes")
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
