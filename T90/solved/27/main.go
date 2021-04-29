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
	m := make(map[string]int)
	var ans []int
	for i := 0; i < n; i++ {
		s[i] = nextString()
		if m[s[i]] == 0 {
			m[s[i]] = i + 1
			ans = append(ans, i+1)
		}
	}
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	for i := range ans {
		fmt.Fprintln(out, ans[i])
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
