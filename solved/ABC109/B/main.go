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
	w := make([]string, n)
	m := make(map[string]bool)
	for i := 0; i < n; i++ {
		w[i] = nextString()
		if m[w[i]] {
			// すでに使ったワード
			fmt.Println("No")
			return
		}
		m[w[i]] = true
	}
	for i := 1; i < n; i++ {
		if w[i][0] != w[i-1][len(w[i-1])-1] {
			// 前の末尾のワードで始まっていない
			fmt.Println("No")
			return
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
