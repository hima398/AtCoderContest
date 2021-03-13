package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func SolveCommentary(n int) {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var dfs func(s string, mc byte)
	dfs = func(s string, mc byte) {
		if len(s) == n {
			fmt.Fprintln(out, s)
		} else {
			for c := byte('a'); c <= mc; c++ {
				ns := s + string(c)
				var nc byte
				if c == mc {
					nc = mc + byte(1)
				} else {
					nc = mc
				}
				dfs(ns, nc)
			}
		}
	}

	dfs("", byte('a'))
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	SolveCommentary(n)
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
