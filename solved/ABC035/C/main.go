package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, bufio.MaxScanTokenSize)
	sc.Split(bufio.ScanWords)

	n, q := nextInt(), nextInt()
	l, r := make([]int, q), make([]int, q)
	for i := 0; i < q; i++ {
		l[i], r[i] = nextInt(), nextInt()
	}
	s := make([]int, n+1)
	for i := 0; i < q; i++ {
		s[l[i]-1]++
		s[r[i]]--
	}
	for i := 0; i < n; i++ {
		s[i+1] += s[i]
	}
	ans := make([]string, n)
	for i := 0; i < n; i++ {
		ans[i] = string(s[i]%2 + '0')
	}
	fmt.Println(strings.Join(ans, ""))
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
